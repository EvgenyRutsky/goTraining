package main

import (
	"fmt"
	"httpserver/apiserver"
	"httpserver/apiserver/handlers"
	"httpserver/infrastructure"
	"httpserver/infrastructure/client"
)

func main() {
	cfg := client.NewConfig()
	bookRepo := infrastructure.NewBookRepository(cfg)
	readerRepo := infrastructure.NewReaderRepository(cfg)

	bookHandler := handlers.NewBookHandler(bookRepo)
	readerHandler := handlers.NewReaderHandler(readerRepo)

	server := apiserver.NewServer(":8080", bookHandler,readerHandler)
	server.ConfigureAndRun()
	fmt.Println("starting server and port 8080")
}
