package main

import (
	"log"
	"net/http"

	"bitbucket.org/strider2038/event-router/requestHandling"
	"github.com/gorilla/mux"
)

func main() {
	handler := requestHandling.NewMessageCollectionRequestHandler()

	router := mux.NewRouter()
	router.HandleFunc("/", handler.HandleRequest).Methods("POST")

	http.Handle("/", router)

	log.Println("Starting server...")
	err := http.ListenAndServe(":3000", nil)
	log.Fatal(err)
}
