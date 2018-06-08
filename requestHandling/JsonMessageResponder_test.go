package requestHandling

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestWriteResponse_statusAndTitle_jsonResponseCreatedAndWritten(t *testing.T) {
	responder := JsonMessageResponder{}
	writer := httptest.NewRecorder()

	result := responder.WriteResponse(writer, http.StatusBadRequest, "message contents")
	response := writer.Result()
	responseBody, _ := ioutil.ReadAll(response.Body)
	contentType := response.Header.Get("Content-Type")

	assert.Nil(t, result)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.Contains(t, string(responseBody), "message")
	message := gjson.Get(string(responseBody), "message").String()
	assert.Contains(t, message, "message contents")
	assert.Contains(t, contentType, "application/json")
}
