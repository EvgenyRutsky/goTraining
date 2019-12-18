package main

import (
	"fmt"
	"server/apiserver"
	"server/apiserver/handlers"
	"server/infrastructure"
)

func main() {
	bookRepo := infrastructure.NewBookRepository()
	readerRepo := infrastructure.NewReaderRepository()

	bookHandler := handlers.NewBookHandler(bookRepo)
	readerHandler := handlers.NewReaderHandler(readerRepo)

	server := apiserver.NewServer(":8080", bookHandler,readerHandler)
	server.ConfigureAndRun()
	fmt.Println("starting server and port 8080")
}
