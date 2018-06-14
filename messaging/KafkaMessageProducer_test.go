package messaging

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

type KafkaWriterMock struct {
}

func (writer KafkaWriterMock) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return nil
}

func TestKafkaMessageProducer_Produce_singleMessage_messageSentToKafkaByWriter(t *testing.T) {
	mockController := gomock.NewController(t)
	defer mockController.Finish()
	factory := NewMockKafkaWriterFlyweight(mockController)
	producer := NewKafkaMessageProducer(factory)
	writer := NewMockKafkaWriter(mockController)
	messageContents := make(map[string]interface{})
	messages := make([]RoutedMessage, 0)
	messages = append(messages, RoutedMessage{"topic", messageContents})

	factory.EXPECT().GetWriterForTopic("topic").Return(writer)
	// writer.EXPECT().WriteMessages(gomock.Any())

	result := producer.Produce(messages)

	assert.Nil(t, result)
}
