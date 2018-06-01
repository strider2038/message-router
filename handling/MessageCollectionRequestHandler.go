package handling

import (
	"net/http"
)

type MessageCollectionRequestHandler struct {
}

func (handler MessageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "Request body cannot be empty", http.StatusBadRequest)
}
