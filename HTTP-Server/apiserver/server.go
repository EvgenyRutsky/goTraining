package apiserver

import (
	"fmt"
	"net/http"
	"server/apiserver/handlers"
)

type Server struct {
	port string
	bookHandler handlers.BookHandler
	readerHandler handlers.ReaderHandler
}

func NewServer (port string, bookHandler handlers.BookHandler, readerHandler handlers.ReaderHandler) *Server {
	return &Server{
		port:        port,
		bookHandler: bookHandler,
		readerHandler: readerHandler,
	}
}

func (s *Server) ConfigureAndRun() {
	fmt.Println("configuring new server")
	bookMux := http.NewServeMux()
	bookMux.HandleFunc("/books/addbook", s.bookHandler.AddBook)
	bookMux.HandleFunc("/books/updatebook", s.bookHandler.UpdateBook)
	bookMux.HandleFunc("/books/deletebook", s.bookHandler.DeleteBook)
	bookMux.HandleFunc("/books/getbook", s.bookHandler.GetBookByID)
	bookMux.HandleFunc("/books/getbooks", s.bookHandler.GetBooks)

	readerMux := http.NewServeMux()
	readerMux.HandleFunc("/readers/addreader", s.readerHandler.AddReader)
	readerMux.HandleFunc("/readers/updatereader", s.readerHandler.UpdateReader)
	readerMux.HandleFunc("/readers/deletereader", s.readerHandler.DeleteReader)
	readerMux.HandleFunc("/readers/getreader", s.readerHandler.GetReaderByID)
	readerMux.HandleFunc("/readers/getreaders", s.readerHandler.GetReaders)

	siteMux := http.NewServeMux()
	siteMux.Handle("/books/", bookMux)
	siteMux.Handle("/readers/", readerMux)
	http.ListenAndServe(s.port, siteMux)
	fmt.Println("listening 8080")
}