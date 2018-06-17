package messaging

import (
	"context"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockKafkaWriter struct {
	mock.Mock
}

func (mock MockKafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	args := mock.Called()

	return args.Error(0)
}

type MockKafkaWriterFlyweight struct {
	mock.Mock
}

func (mock *MockKafkaWriterFlyweight) GetWriterForTopic(topicName string) KafkaWriter {
	args := mock.Called(topicName)

	return args.Get(0).(KafkaWriter)
}

func (mock *MockKafkaWriterFlyweight) PoolSize() int {
	panic("not implemented")
}

func TestKafkaMessageProducer_Produce_singleMessage_messageSentToKafkaByWriter(t *testing.T) {
	writer := &MockKafkaWriter{}
	factory := MockKafkaWriterFlyweight{}
	producer := NewKafkaMessageProducer(&factory)
	messageContents := make(map[string]interface{})
	messages := make([]RoutedMessage, 0)
	messages = append(messages, RoutedMessage{"topic", messageContents})
	factory.On("GetWriterForTopic", "topic").Return(writer)
	writer.On("WriteMessages").Return(nil)

	result := producer.Produce(messages)

	factory.Mock.AssertExpectations(t)
	assert.Nil(t, result)
}
