package messaging

import (
	"bitbucket.org/strider2038/event-router/data"
	"github.com/stretchr/testify/mock"
)

type MessageProducerMock struct {
	mock.Mock
}

func (mock *MessageProducerMock) Produce(message *data.MessagePack) error {
	arguments := mock.Called(message)

	return arguments.Error(0)
}
