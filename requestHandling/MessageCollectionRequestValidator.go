package requestHandling

import (
	"net/http"
	"regexp"
)

type MessageCollectionRequestValidator struct {
}

func (validator MessageCollectionRequestValidator) ValidateRequest(request *http.Request) *RequestError {
	var requestError *RequestError = nil

	if request.Body == nil {
		requestError = &RequestError{http.StatusBadRequest, "Request body cannot be empty"}
	} else if request.Method != "POST" || request.URL.Path != "/" {
		requestError = &RequestError{http.StatusNotFound, "Route not found"}
	} else {
		contentType := request.Header.Get("Content-Type")
		contentTypeIsJson, _ := regexp.MatchString("application/json.*", contentType)
		if !contentTypeIsJson {
			requestError = &RequestError{http.StatusBadRequest, "Content type must be application/json"}
		}
	}

	return requestError
}
