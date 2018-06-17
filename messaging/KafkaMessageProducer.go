package messaging

import (
	"context"
	"encoding/json"

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
		kafkaMessage := kafka.Message{}
		contents, _ := json.Marshal(message.Message)
		kafkaMessage.Value = contents
		writer.WriteMessages(context.Background(), kafkaMessage)
	}

	return nil
}
