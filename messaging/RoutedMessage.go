package messaging

type RoutedMessage struct {
	Topic   string                 `json:"topic"`
	Message map[string]interface{} `json:"message"`
}
