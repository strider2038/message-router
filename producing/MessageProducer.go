package producing

import (
	"github.com/strider2038/message-router/data"
)

type MessageProducer interface {
	Produce(message *data.MessagePack) error
}
