package requestHandling

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type ResponderMock struct {
	mock.Mock
}

func (mock *ResponderMock) WriteResponse(writer http.ResponseWriter, status int, message string) error {
	arguments := mock.Called(writer, status, message)

	return arguments.Error(0)
}
