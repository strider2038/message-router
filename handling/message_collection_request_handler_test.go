package handling

import (
	"testing"

	mockHttp "bitbucket.org/strider2038/event-router/mocks/http"
	"github.com/golang/mock/gomock"
	"net/http"
)

func TestHandleRequest_emptyBody_httpErrorReturned(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	writer := mockHttp.NewMockResponseWriter(mockController)
	handler := MessageCollectionRequestHandler{}
	request := http.Request{}

	writer.EXPECT().Write([]byte("Hello"))

	handler.HandleRequest(writer, &request)
}
