package messaging

type KafkaWriterFlyweight interface {
	GetWriterForTopic(topicName string) KafkaWriter
	PoolSize() int
}
