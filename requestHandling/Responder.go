package requestHandling

import "net/http"

type Responder interface {
	WriteResponse(writer http.ResponseWriter, status int, message string) error
}
