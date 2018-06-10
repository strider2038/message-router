package messaging

import "github.com/segmentio/kafka-go"

type KafkaWriterFlyweight interface {
	GetWriterForTopic(topicName string) *kafka.Writer
	PoolSize() int
}
