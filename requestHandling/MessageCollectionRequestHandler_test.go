package requestHandling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
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
	assert.Contains(t, string(responseBody), "title")
	title := gjson.Get(string(responseBody), "title").String()
	assert.Contains(t, title, "Request body cannot be empty")
	assert.Contains(t, contentType, "application/json")
}
