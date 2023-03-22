package api

import (
	"log"
	"net/http"
)

func RunServer() {
	log.Printf("Starting server on port 4000..\n")
	err := http.ListenAndServe(":4000", router())
	if err != nil {
		log.Fatal(err)
	}
}
