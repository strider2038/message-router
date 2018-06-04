package requestHandling

import "net/http"

type RequestHandler interface {
	HandleRequest(writer http.ResponseWriter, request *http.Request)
}