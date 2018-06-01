package main

import (
	"log"
	"net/http"

	"bitbucket.org/strider2038/event-router/requestHandling"
)

func main() {
	handler := requestHandling.MessageCollectionRequestHandler{}
	http.HandleFunc("/", handler.HandleRequest)

	log.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
