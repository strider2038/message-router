package producing

import (
	"context"
	"encoding/json"

	"bitbucket.org/strider2038/event-router/data"
	"github.com/segmentio/kafka-go"
)

type kafkaMessageProducer struct {
	writerFactory KafkaWriterFactory
}

func NewKafkaMessageProducer(writerFactory KafkaWriterFactory) *kafkaMessageProducer {
	return &kafkaMessageProducer{writerFactory}
}

func (producer *kafkaMessageProducer) Produce(message *data.MessagePack) error {
	kafkaMessage := kafka.Message{}
	messageKey, _ := json.Marshal(message.Message.Key)
	messageValue, _ := json.Marshal(message.Message.Value)
	kafkaMessage.Key = messageKey
	kafkaMessage.Value = messageValue

	writer := producer.writerFactory.CreateWriterForTopic(message.Topic)

	return writer.WriteMessages(context.Background(), kafkaMessage)
}
