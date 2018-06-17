package messaging

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type kafkaMessageProducer struct {
	factory KafkaWriterFlyweight
}

func NewKafkaMessageProducer(factory KafkaWriterFlyweight) *kafkaMessageProducer {
	return &kafkaMessageProducer{factory}
}

func (producer *kafkaMessageProducer) Produce(messages []RoutedMessage) error {
	for _, message := range messages {
		writer := producer.factory.GetWriterForTopic(message.Topic)
		writer.WriteMessages(context.Background(), kafka.Message{})
	}

	return nil
}
