package requestHandling

import (
	"encoding/json"
	"io"
	"net/http"
)

type messageCollectionRequestHandler struct {
	validator RequestValidator
}

func NewMessageCollectionRequestHandler(validator RequestValidator) *messageCollectionRequestHandler {
	return &messageCollectionRequestHandler{validator}
}

func (handler messageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	requestError := handler.validator.ValidateRequest(request)

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	if requestError == nil {
		writer.WriteHeader(http.StatusOK)
		messageTitle := "All messages were successfully sent to queue"
		message, _ := json.Marshal(JsonResponse{messageTitle})
		io.WriteString(writer, string(message))
	} else {
		writer.WriteHeader(requestError.status)
		message, _ := json.Marshal(JsonResponse{requestError.title})
		io.WriteString(writer, string(message))
	}
}
