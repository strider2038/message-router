package handling

import (
	"fmt"
	"net/http"
)

type MessageCollectionRequestHandler struct {
}

func (handler MessageCollectionRequestHandler) HandleRequest(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello")
}
