package main

import (
	"context"
	"log"
	"os"

	"github.com/digital-wallet-svc/internal/http"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	err_serv := http.StartServer(ctx, port)

	if err_serv != nil {
		log.Fatal(err_serv)
	}
}
