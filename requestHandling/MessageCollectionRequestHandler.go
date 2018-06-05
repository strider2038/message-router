package requestHandling

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
)

type MessageCollectionRequestHandler struct {
}

func (handler MessageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	contentType := request.Header.Get("Content-Type")
	contentTypeIsJson, _ := regexp.MatchString("application/json.*", contentType)

	var messageTitle string

	if request.Body == nil {
		writer.WriteHeader(http.StatusBadRequest)
		messageTitle = "Request body cannot be empty"
	} else if !contentTypeIsJson {
		writer.WriteHeader(http.StatusBadRequest)
		messageTitle = "Content type must be application/json"
	} else if request.Method != "POST" || request.URL.Path != "/" {
		writer.WriteHeader(http.StatusNotFound)
		messageTitle = "Route not found"
	} else {
		writer.WriteHeader(http.StatusOK)
		messageTitle = "All messages were successfully sent to queue"
	}

	message, _ := json.Marshal(JsonResponse{messageTitle})
	io.WriteString(writer, string(message))
}
