package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":8080"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("The server is running on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start http server")
	}
}
