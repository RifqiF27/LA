package main

import (
	"cinema/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}