package main

import (
	"log"
	"net/http"
	common "oms-common"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr          = common.EnvString("HTTP_ADDR", ":8080")
	ordersServiceAddr = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(ordersServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at: %v", ordersServiceAddr)

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("The server is running on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("failed to start http server")
	}
}
