package requestHandling

import (
	"encoding/json"
	"io"
	"net/http"
)

type MessageCollectionRequestHandler struct {
}

func (handler MessageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusBadRequest)
	message, _ := json.Marshal(JsonResponse{"Request body cannot be empty"})
	io.WriteString(writer, string(message))
}
