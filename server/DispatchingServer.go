package server

import (
	"fmt"
	"github.com/bitwurx/jrpc2"
	"github.com/segmentio/kafka-go"
	"github.com/strider2038/message-router/config"
	"github.com/strider2038/message-router/producing"
)

// HTTP server for handling messages received by HTTP protocol (JSON RPC v2)
type DispatchingServer interface {
	Start()
}

// Dispatching server constructor
func NewDispatchingServer(config config.Config) DispatchingServer {
	var headers map[string]string

	writerConfig := kafka.WriterConfig{
		Brokers:  config.Brokers,
		Balancer: &kafka.LeastBytes{},
	}

	factory := producing.NewKafkaWriterFlyweightFactory(writerConfig)
	producer := producing.NewKafkaMessageProducer(factory)
	dispatcher := messageDispatcher{producer}

	host := fmt.Sprintf(":%d", config.Port)
	server := jrpc2.NewServer(host, "/rpc", headers)
	server.Register("dispatch", jrpc2.Method{Method: dispatcher.Handle})

	return server
}
