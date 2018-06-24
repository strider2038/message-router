package server

import (
	"encoding/json"
	"errors"
	"testing"

	"bitbucket.org/strider2038/event-router/producing"
	"github.com/bitwurx/jrpc2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const invalidJsonRequest = `{`
const emptyParamsBody = `{}`
const noMessageBody = `
{
	"topic": "topic"
}
`
const emptyMessageBody = `
{
	"topic": "topic",
	"message": {}
}
`
const validRequestBody = `
{
	"topic": "topic",
	"message": {
		"value": {
			"valid": true
		}
	}
}
`

func TestMessageDispatcher_Handle_InvalidRequestBody_ErrorReturned(t *testing.T) {
	dispatcher := messageDispatcher{}

	_, err := dispatcher.Handle([]byte(invalidJsonRequest))

	assert.Equal(t, jrpc2.InvalidParamsCode, err.Code)
}

func TestMessageDispatcher_Handle_InvalidParamsBody_ErrorReturned(t *testing.T) {
	tests := []struct {
		name   string
		params string
	}{
		{
			"empty params",
			emptyParamsBody,
		},
		{
			"no message",
			noMessageBody,
		},
		{
			"empty message",
			emptyMessageBody,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			dispatcher := messageDispatcher{}

			_, err := dispatcher.Handle(json.RawMessage(test.params))

			assert.Equal(t, jrpc2.InvalidParamsCode, err.Code)
		})
	}
}

func TestMessageDispatcher_Handle_ValidMessage_MessageProducedAndSuccessResultReturned(t *testing.T) {
	producer := producing.MessageProducerMock{}
	dispatcher := messageDispatcher{&producer}
	producer.On("Produce", mock.Anything).Return(nil)

	result, err := dispatcher.Handle(json.RawMessage(validRequestBody))

	assert.Nil(t, err)
	assert.Equal(t, "Message was successfully dispatched to topic \"topic\".", result)
	producer.AssertExpectations(t)
}

func TestMessageDispatcher_Handle_ValidMessage_ProducingFailedAndErrorReturned(t *testing.T) {
	producer := producing.MessageProducerMock{}
	dispatcher := messageDispatcher{&producer}
	producer.On("Produce", mock.Anything).Return(errors.New("producing failed"))

	_, err := dispatcher.Handle(json.RawMessage(validRequestBody))

	producer.AssertExpectations(t)
	assert.Equal(t, jrpc2.InternalErrorCode, err.Code)
}
