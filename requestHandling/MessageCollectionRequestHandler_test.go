package requestHandling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestHandleRequest_emptyBody_badRequestErrorReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}
	request := httptest.NewRequest("POST", "/", nil)
	writer := httptest.NewRecorder()

	handler.HandleRequest(writer, request)
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(
		t,
		"{\"title\":\"Request body cannot be empty\"}",
		string(responseBody))
	assert.Contains(t, contentType, "application/json")
}
