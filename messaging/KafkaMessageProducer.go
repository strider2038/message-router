package messaging

import (
	"context"
	"encoding/json"

	"bitbucket.org/strider2038/event-router/data"
	"github.com/segmentio/kafka-go"
)

type kafkaMessageProducer struct {
	factory KafkaWriterFlyweight
}

func NewKafkaMessageProducer(factory KafkaWriterFlyweight) *kafkaMessageProducer {
	return &kafkaMessageProducer{factory}
}

func (producer *kafkaMessageProducer) Produce(message *data.MessagePack) error {
	kafkaMessage := kafka.Message{}
	messageKey, _ := json.Marshal(message.Message.Key)
	messageValue, _ := json.Marshal(message.Message.Value)
	kafkaMessage.Key = messageKey
	kafkaMessage.Value = messageValue

	writer := producer.factory.GetWriterForTopic(message.Topic)

	return writer.WriteMessages(context.Background(), kafkaMessage)
}
