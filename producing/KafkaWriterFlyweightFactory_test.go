package producing

import (
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestKafkaWriterFlyweightFactory_CreateWriterForTopic_TopicName_WriterCreatedAndReturnedAndPoolSizeIsOne(t *testing.T) {
	factory := createKafkaWriterFlyweightFactory()

	writer := factory.CreateWriterForTopic("topic")

	assert.NotNil(t, writer)
	assert.Equal(t, 1, factory.PoolSize())
}

func TestKafkaWriterFlyweightFactory_CreateWriterForTopic_TwoEqualTopicsAndOneDifferent_PoolSizeIsThree(t *testing.T) {
	factory := createKafkaWriterFlyweightFactory()

	factory.CreateWriterForTopic("topic")
	factory.CreateWriterForTopic("topic")
	factory.CreateWriterForTopic("another-topic")

	assert.Equal(t, 2, factory.PoolSize())
}

func createKafkaWriterFlyweightFactory() *kafkaWriterFlyweightFactory {
	return NewKafkaWriterFlyweightFactory(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
	})
}
