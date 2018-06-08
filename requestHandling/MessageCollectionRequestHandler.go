package requestHandling

import (
	"net/http"
)

type messageCollectionRequestHandler struct {
	responder Responder
}

func NewMessageCollectionRequestHandler(responder Responder) *messageCollectionRequestHandler {
	return &messageCollectionRequestHandler{responder}
}

func (handler messageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	handler.responder.WriteResponse(writer, http.StatusOK, "All messages were successfully sent to queue")
}
