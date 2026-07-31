package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	if err != nil {
	}
}
