package handling

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func TestHandleRequest_emptyBody_httpErrorReturned(t *testing.T) {
	handler := MessageCollectionRequestHandler{}
	request := httptest.NewRequest("POST", "/", nil)
	writer := httptest.NewRecorder()

	handler.HandleRequest(writer, request)
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Equal(t, "Request body cannot be empty\n", string(responseBody))
}
