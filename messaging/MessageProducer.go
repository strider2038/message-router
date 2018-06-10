package messaging

type MessageProducer interface {
	Produce(messages []RoutedMessage) error
}
