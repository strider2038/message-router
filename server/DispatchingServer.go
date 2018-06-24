package server

import (
	"github.com/bitwurx/jrpc2"
)

type DispatchingServer interface {
	Start()
}

func NewDispatchingServer() DispatchingServer {
	var headers map[string]string

	dispatcher := messageDispatcher{}

	server := jrpc2.NewServer(":3000", "/rpc", headers)
	server.Register("dispatch", jrpc2.Method{Method: dispatcher.Handle})

	return server
}
