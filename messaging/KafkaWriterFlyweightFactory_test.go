package messaging

import (
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestKafkaWriterFlyweightFactory_GetWriterForTopic_topicName_writerCreatedAndReturnedAndPoolSizeIsOne(t *testing.T) {
	factory := createKafkaWriterFlyweightFactory()

	writer := factory.GetWriterForTopic("topic")

	assert.NotNil(t, writer)
	assert.Equal(t, 1, factory.PoolSize())
}

func TestKafkaWriterFlyweightFactory_GetWriterForTopic_twoEqualTopicsAndOneDifferent_poolSizeIsThree(t *testing.T) {
	factory := createKafkaWriterFlyweightFactory()

	factory.GetWriterForTopic("topic")
	factory.GetWriterForTopic("topic")
	factory.GetWriterForTopic("another-topic")

	assert.Equal(t, 2, factory.PoolSize())
}

func createKafkaWriterFlyweightFactory() KafkaWriterFlyweight {
	return NewKafkaWriterFlyweightFactory(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
	})
}
