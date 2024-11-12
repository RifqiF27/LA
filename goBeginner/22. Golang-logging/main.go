package main

import (
	"book-store/wire"
	"log"
	"net/http"
)

func main() {
	r, err := wire.InitializeRouterHandler()
	if err != nil {
		log.Fatalf("failed to initialize router: %v", err)
	}
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}