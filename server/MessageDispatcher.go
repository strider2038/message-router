package server

import (
	"encoding/json"
	"fmt"

	"bitbucket.org/strider2038/event-router/data"
	"bitbucket.org/strider2038/event-router/producing"
	"github.com/asaskevich/govalidator"
	"github.com/bitwurx/jrpc2"
)

type messageDispatcher struct {
	producer producing.MessageProducer
}

func (dispatcher *messageDispatcher) Handle(params json.RawMessage) (interface{}, *jrpc2.ErrorObject) {
	pack := new(data.MessagePack)

	if err := jrpc2.ParseParams(params, pack); err != nil {
		return nil, err
	}

	_, err := govalidator.ValidateStruct(pack)
	if err != nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InvalidParamsCode,
			Message: jrpc2.InvalidParamsMsg,
			Data:    err.Error(),
		}
	}

	err = dispatcher.producer.Produce(pack)

	if err != nil {
		return nil, &jrpc2.ErrorObject{
			Code:    jrpc2.InternalErrorCode,
			Message: jrpc2.InternalErrorMsg,
			Data:    err.Error(),
		}
	}

	return fmt.Sprintf("Message was successfully dispatched to topic \"%s\".", pack.Topic), nil
}
