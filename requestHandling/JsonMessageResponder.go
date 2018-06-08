package requestHandling

import (
	"encoding/json"
	"io"
	"net/http"
)

type JsonMessageResponder struct {
}

func (responder JsonMessageResponder) WriteResponse(writer http.ResponseWriter, status int, message string) error {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(status)
	response, _ := json.Marshal(JsonResponse{message})
	io.WriteString(writer, string(response))

	return nil
}
