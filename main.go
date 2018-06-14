package main

import (
	"log"
	"net/http"

	"bitbucket.org/strider2038/event-router/requestHandling"
	"github.com/gorilla/mux"
)

func main() {
	responder := requestHandling.JsonMessageResponder{}
	handler := requestHandling.NewMessageCollectionRequestHandler(responder)

	router := mux.NewRouter()
	router.HandleFunc("/", handler.HandleRequest).Methods("POST")

	http.Handle("/", router)

	log.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)

	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{"localhost:9092"},
	// 	Topic:   "topic-A",
	// 	Balancer: &kafka.LeastBytes{},
	// })
	//
	// err := w.WriteMessages(context.Background(),
	// 	kafka.Message{
	// 		Key:   []byte("Key-A"),
	// 		Value: []byte("Hello World!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-B"),
	// 		Value: []byte("One!"),
	// 	},
	// 	kafka.Message{
	// 		Key:   []byte("Key-C"),
	// 		Value: []byte("Two!"),
	// 	},
	// )
	//
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// w.Close()
}
