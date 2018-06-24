package server

import (
	"encoding/json"

	"github.com/bitwurx/jrpc2"
)

type RemoteProcedure interface {
	Handle(params json.RawMessage) (interface{}, *jrpc2.ErrorObject)
}
