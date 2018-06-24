package server

import (
	"bitbucket.org/strider2038/event-router/messaging"
	"github.com/bitwurx/jrpc2"
	"github.com/segmentio/kafka-go"
)

type DispatchingServer interface {
	Start()
}

func NewDispatchingServer() DispatchingServer {
	var headers map[string]string

	config := kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Balancer: &kafka.LeastBytes{},
	}
	factory := messaging.NewKafkaWriterFlyweightFactory(config)
	producer := messaging.NewKafkaMessageProducer(factory)
	dispatcher := messageDispatcher{producer}

	server := jrpc2.NewServer(":3000", "/rpc", headers)
	server.Register("dispatch", jrpc2.Method{Method: dispatcher.Handle})

	return server
}
