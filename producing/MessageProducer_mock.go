package producing

import (
	"github.com/stretchr/testify/mock"
	"github.com/strider2038/message-router/data"
)

type MessageProducerMock struct {
	mock.Mock
}

func (mock *MessageProducerMock) Produce(message *data.MessagePack) error {
	arguments := mock.Called(message)

	return arguments.Error(0)
}
