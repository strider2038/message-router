package requestHandling

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"

	"bitbucket.org/strider2038/event-router/messaging"
	"github.com/stretchr/testify/mock"
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
{
	"something": "invalid"
}
`

func TestMessageCollectionRequestHandler_HandleRequest_validRequest_okReturned(t *testing.T) {
	responder := &ResponderMock{}
	producer := &messaging.MessageProducerMock{}
	handler := NewMessageCollectionRequestHandler(responder, producer)
	bodyReader := strings.NewReader(validRequestBody)
	request := httptest.NewRequest("POST", "/", bodyReader)
	writer := httptest.NewRecorder()
	producer.On("Produce", mock.Anything).Return(nil)
	responder.On("WriteResponse", writer, http.StatusOK, "All messages were successfully sent to queue").Return(nil)

	handler.HandleRequest(writer, request)

	producer.AssertExpectations(t)
	responder.AssertExpectations(t)
}

func TestMessageCollectionRequestHandler_HandleRequest_invalidRequest_okReturned(t *testing.T) {
	responder := &ResponderMock{}
	producer := &messaging.MessageProducerMock{}
	handler := NewMessageCollectionRequestHandler(responder, producer)
	bodyReader := strings.NewReader(invalidRequestBody)
	request := httptest.NewRequest("POST", "/", bodyReader)
	writer := httptest.NewRecorder()
	responder.
		On(
			"WriteResponse",
			writer,
			http.StatusBadRequest,
			mock.Anything).
		Return(nil)

	handler.HandleRequest(writer, request)

	responder.AssertExpectations(t)
	producer.AssertNotCalled(t, "Produce")
}
