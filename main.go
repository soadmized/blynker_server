package main

import (
	"blynker/internal/api"
	"log"
	"net/http"
)

func main() {
	server := api.New()

	err := http.ListenAndServe(":8080", server.Router)
	if err != nil {
		log.Fatal(err)
	}
}