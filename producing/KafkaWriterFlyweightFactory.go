package producing

import "github.com/segmentio/kafka-go"

type kafkaWriterFlyweightFactory struct {
	pool   map[string]kafka.Writer
	config kafka.WriterConfig
}

func NewKafkaWriterFlyweightFactory(config kafka.WriterConfig) *kafkaWriterFlyweightFactory {
	pool := make(map[string]kafka.Writer)

	return &kafkaWriterFlyweightFactory{pool, config}
}

func (factory *kafkaWriterFlyweightFactory) CreateWriterForTopic(topicName string) KafkaWriter {
	if writer, exists := factory.pool[topicName]; exists {
		return &writer
	}

	factory.config.Topic = topicName
	writer := kafka.NewWriter(factory.config)

	factory.pool[topicName] = *writer

	return writer
}

func (factory *kafkaWriterFlyweightFactory) PoolSize() int {
	return len(factory.pool)
}
