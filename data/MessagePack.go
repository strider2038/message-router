package data

import "errors"

// Apache Kafka message structure
type Message struct {
	Key   map[string]interface{} `json:"key"   valid:"-"`
	Value map[string]interface{} `json:"value" valid:"-"`
}

// Message pack includes topic for kafka and message value
type MessagePack struct {
	Topic   string  `json:"topic"   valid:"required"`
	Message Message `json:"message" valid:"required"`
}

// Special method for JSON RPC v2 library
func (pack *MessagePack) FromPositional(params []interface{}) error {
	return errors.New("invalid message pack format")
}
