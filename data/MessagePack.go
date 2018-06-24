package data

import "errors"

type Message struct {
	Key   map[string]interface{} `json:"key"   valid:"-"`
	Value map[string]interface{} `json:"value" valid:"-"`
}

type MessagePack struct {
	Topic   string  `json:"topic"   valid:"required"`
	Message Message `json:"message" valid:"required"`
}

func (pack *MessagePack) FromPositional(params []interface{}) error {
	return errors.New("invalid message pack format")
}
