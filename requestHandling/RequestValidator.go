package requestHandling

import "net/http"

type RequestError struct {
	status int
	title  string
}

type RequestValidator interface {
	ValidateRequest(request *http.Request) *RequestError
}
