package producing

import (
	"context"
	"errors"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/strider2038/message-router/data"
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

func (mock *MockKafkaWriterFlyweight) CreateWriterForTopic(topicName string) KafkaWriter {
	args := mock.Called(topicName)

	return args.Get(0).(KafkaWriter)
}

func (mock *MockKafkaWriterFlyweight) PoolSize() int {
	panic("not implemented")
}

func TestKafkaMessageProducer_Produce_Message_MessageSentToKafkaByWriter(t *testing.T) {
	writer := &MockKafkaWriter{}
	factory := MockKafkaWriterFlyweight{}
	producer := NewKafkaMessageProducer(&factory)
	message := data.MessagePack{"topic", data.Message{}}
	factory.On("CreateWriterForTopic", "topic").Return(writer)
	writer.On("WriteMessages").Return(nil)

	err := producer.Produce(&message)

	factory.Mock.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestKafkaMessageProducer_Produce_Message_MessageSentFailedAndErrorReturned(t *testing.T) {
	writer := &MockKafkaWriter{}
	factory := MockKafkaWriterFlyweight{}
	producer := NewKafkaMessageProducer(&factory)
	message := data.MessagePack{"topic", data.Message{}}
	factory.On("CreateWriterForTopic", "topic").Return(writer)
	writer.On("WriteMessages").Return(errors.New("error"))

	err := producer.Produce(&message)

	factory.Mock.AssertExpectations(t)
	assert.Equal(t, "error", err.Error())
}
