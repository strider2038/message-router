package messaging

import "github.com/stretchr/testify/mock"

type MessageProducerMock struct {
	mock.Mock
}

func (mock *MessageProducerMock) Produce(messages []RoutedMessage) error {
	arguments := mock.Called(messages)

	return arguments.Error(0)
}
