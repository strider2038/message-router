package requestHandling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

func TestHandleRequest_emptyBody_badRequestErrorReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}
	request := httptest.NewRequest("POST", "/", nil)
	request.Body = nil
	writer := httptest.NewRecorder()

	handler.HandleRequest(writer, request)
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Contains(t, string(responseBody), "title")
	title := gjson.Get(string(responseBody), "title").String()
	assert.Contains(t, title, "Request body cannot be empty")
	assert.Contains(t, contentType, "application/json")
}

var invalidRouteCases = []struct {
	method string
	url    string
}{
	{"POST", "/url"},
	{"GET", "/"},
	{"PUT", "/"},
}

func TestHandleRequest_invalidRoute_notFoundErrorReturned(t *testing.T) {
	for _, testCase := range invalidRouteCases {
		handler := MessageCollectionRequestHandler{}
		bodyReader := strings.NewReader("body")
		request := httptest.NewRequest(testCase.method, testCase.url, bodyReader)
		request.Header.Set("Content-Type", "application/json; charset=utf-8")
		writer := httptest.NewRecorder()

		handler.HandleRequest(writer, request)
		response := writer.Result()
		responseBody, _ := ioutil.ReadAll(response.Body)
		contentType := response.Header.Get("Content-Type")

		assert.Equal(t, http.StatusNotFound, response.StatusCode)
		assert.Contains(t, string(responseBody), "title")
		title := gjson.Get(string(responseBody), "title").String()
		assert.Contains(t, title, "Route not found")
		assert.Contains(t, contentType, "application/json")
	}
}

func TestHandleRequest_contentTypeIsNotJson_badRequestErrorReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}
	bodyReader := strings.NewReader("body")
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "text/html; charset=utf-8")
	writer := httptest.NewRecorder()

	handler.HandleRequest(writer, request)
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Contains(t, string(responseBody), "title")
	title := gjson.Get(string(responseBody), "title").String()
	assert.Contains(t, title, "Content type must be application/json")
	assert.Contains(t, contentType, "application/json")
}

func TestHandleRequest_validRequest_okReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}
	bodyReader := strings.NewReader("body")
	request := httptest.NewRequest("POST", "/", bodyReader)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	writer := httptest.NewRecorder()

	handler.HandleRequest(writer, request)
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Contains(t, string(responseBody), "title")
	title := gjson.Get(string(responseBody), "title").String()
	assert.Contains(t, title, "All messages were successfully sent to queue")
	assert.Contains(t, contentType, "application/json")
}
