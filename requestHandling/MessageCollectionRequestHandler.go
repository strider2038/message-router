package requestHandling

import (
	"encoding/json"
	"io"
	"net/http"
)

type messageCollectionRequestHandler struct {
}

func NewMessageCollectionRequestHandler() *messageCollectionRequestHandler {
	return &messageCollectionRequestHandler{}
}

func (handler messageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)
	messageTitle := "All messages were successfully sent to queue"
	message, _ := json.Marshal(JsonResponse{messageTitle})
	io.WriteString(writer, string(message))
}
