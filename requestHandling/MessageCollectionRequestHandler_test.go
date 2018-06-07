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

func TestHandleRequest_validRequest_okReturned(t *testing.T) {
	handler := NewMessageCollectionRequestHandler()
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
