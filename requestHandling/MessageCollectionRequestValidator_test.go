package requestHandling

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateRequest_validRequest_nilReturned(t *testing.T) {
	validator := MessageCollectionRequestValidator{}
	bodyReader := strings.NewReader("body")
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	requestError := validator.ValidateRequest(request)

	assert.Nil(t, requestError)
}

func TestValidateRequest_emptyBody_badRequestErrorReturned(t *testing.T) {
	validator := MessageCollectionRequestValidator{}
	request := httptest.NewRequest("POST", "/", nil)
	request.Body = nil

	requestError := validator.ValidateRequest(request)

	assert.NotNil(t, requestError)
	assert.Equal(t, http.StatusBadRequest, requestError.status)
	assert.Equal(t, "Request body cannot be empty", requestError.title)
}

var invalidRouteCases = []struct {
	method string
	url    string
}{
	{"POST", "/url"},
	{"GET", "/"},
	{"PUT", "/"},
}

func TestValidateRequest_invalidRoute_notFoundErrorReturned(t *testing.T) {
	for _, testCase := range invalidRouteCases {
		validator := MessageCollectionRequestValidator{}
		bodyReader := strings.NewReader("body")
		request := httptest.NewRequest(testCase.method, testCase.url, bodyReader)
		request.Header.Set("Content-Type", "application/json; charset=utf-8")

		requestError := validator.ValidateRequest(request)

		assert.NotNil(t, requestError)
		assert.Equal(t, http.StatusNotFound, requestError.status)
		assert.Equal(t, "Route not found", requestError.title)
	}
}

func TestValidateRequest_contentTypeIsNotJson_badRequestErrorReturned(t *testing.T) {
	validator := MessageCollectionRequestValidator{}
	bodyReader := strings.NewReader("body")
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "text/html; charset=utf-8")

	requestError := validator.ValidateRequest(request)

	assert.NotNil(t, requestError)
	assert.Equal(t, http.StatusBadRequest, requestError.status)
	assert.Equal(t, "Content type must be application/json", requestError.title)
}
