package messaging

import (
	"bitbucket.org/strider2038/event-router/data"
)

type MessageProducer interface {
	Produce(message *data.MessagePack) error
}
