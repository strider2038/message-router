package server

import (
	"encoding/json"
	"testing"

	"github.com/bitwurx/jrpc2"
	"github.com/stretchr/testify/assert"
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
