package main

import (
	"log"
	"net/http"

	"bitbucket.org/strider2038/event-router/messaging"
	"bitbucket.org/strider2038/event-router/requestHandling"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

func main() {
	responder := requestHandling.JsonMessageResponder{}
	config := kafka.WriterConfig{
		Brokers:  []string{"kafka:9092"},
		Balancer: &kafka.LeastBytes{},
	}
	factory := messaging.NewKafkaWriterFlyweightFactory(config)
	producer := messaging.NewKafkaMessageProducer(factory)
	handler := requestHandling.NewMessageCollectionRequestHandler(responder, producer)

	router := mux.NewRouter()
	router.HandleFunc("/", handler.HandleRequest).Methods("POST")

	http.Handle("/", router)

	log.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
