package main

import "bitbucket.org/strider2038/event-router/server"

func main() {
	dispatchingServer := server.NewDispatchingServer()
	dispatchingServer.Start()

	// responder := requestHandling.JsonMessageResponder{}
	// config := kafka.WriterConfig{
	// 	Brokers:  []string{"kafka:9092"},
	// 	Balancer: &kafka.LeastBytes{},
	// }
	// factory := messaging.NewKafkaWriterFlyweightFactory(config)
	// producer := messaging.NewKafkaMessageProducer(factory)
	// handler := requestHandling.NewMessageCollectionRequestHandler(responder, producer)
	//
	// router := mux.NewRouter()
	// router.HandleFunc("/", handler.HandleRequest).Methods("POST")
	//
	// http.Call("/", router)
	//
	// log.Println("Starting server...")
	// err := http.ListenAndServe(":3000", nil)
	// log.Fatal(err)
}
