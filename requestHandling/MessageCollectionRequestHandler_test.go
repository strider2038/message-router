package requestHandling

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"bitbucket.org/strider2038/event-router/messaging"
	"github.com/golang/mock/gomock"
)

const validRequestBody = `
[
	{
		"topic": "TopicName",
		"message": {
			"property": "value",
			"object": {
				"a": 1
			}
		}
	}
]
`
const invalidRequestBody = `
[
	{
		"something": "invalid"
	}
]
`

func TestHandleRequest_validRequest_okReturned(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	responder := NewMockResponder(mockController)
	producer := messaging.NewMockMessageProducer(mockController)
	handler := NewMessageCollectionRequestHandler(responder, producer)
	bodyReader := strings.NewReader(validRequestBody)
	request := httptest.NewRequest("POST", "/", bodyReader)
	writer := httptest.NewRecorder()

	producer.EXPECT().Produce(gomock.Any())
	responder.EXPECT().WriteResponse(writer, http.StatusOK, "All messages were successfully sent to queue")

	handler.HandleRequest(writer, request)
}

func TestHandleRequest_invalidRequest_okReturned(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	responder := NewMockResponder(mockController)
	producer := messaging.NewMockMessageProducer(mockController)
	handler := NewMessageCollectionRequestHandler(responder, producer)
	bodyReader := strings.NewReader(invalidRequestBody)
	request := httptest.NewRequest("POST", "/", bodyReader)
	writer := httptest.NewRecorder()

	responder.EXPECT().WriteResponse(writer, http.StatusBadRequest, "json: cannot unmarshal array into Go value of type messaging.MessageCollection")

	handler.HandleRequest(writer, request)
}
