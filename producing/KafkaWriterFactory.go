package producing

type KafkaWriterFactory interface {
	CreateWriterForTopic(topicName string) KafkaWriter
}
