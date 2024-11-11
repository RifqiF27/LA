package main

import (
	"book-store/wire"
	"log"
	"net/http"
)

func main() {
	r := wire.InitializeRouterHandler()

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}