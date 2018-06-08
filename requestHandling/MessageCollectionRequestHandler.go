package requestHandling

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/strider2038/event-router/messaging"
)

type messageCollectionRequestHandler struct {
	responder Responder
}

func NewMessageCollectionRequestHandler(responder Responder) *messageCollectionRequestHandler {
	return &messageCollectionRequestHandler{responder}
}

func (handler messageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	messages := make([]messaging.RoutedMessage, 0)
	err := json.NewDecoder(request.Body).Decode(&messages)

	if err != nil {
		handler.responder.WriteResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		handler.responder.WriteResponse(writer, http.StatusOK, "All messages were successfully sent to queue")
	}
}
