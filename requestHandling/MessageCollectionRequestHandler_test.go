package requestHandling

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
)

func TestHandleRequest_validRequest_okReturned(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	responder := NewMockResponder(mockController)
	handler := NewMessageCollectionRequestHandler(responder)
	request := httptest.NewRequest("POST", "/", nil)
	writer := httptest.NewRecorder()

	responder.EXPECT().WriteResponse(writer, http.StatusOK, "All messages were successfully sent to queue")

	handler.HandleRequest(writer, request)
}
