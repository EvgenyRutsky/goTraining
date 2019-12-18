package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/domain"
	"server/infrastructure"
	"strconv"
)

type BookHandler interface {
	AddBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	GetBookByID(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
}

type bookHandler struct {
	bookRepository infrastructure.BooksRepository
}

func NewBookHandler(bookRepository infrastructure.BooksRepository) BookHandler {
	return &bookHandler{
		bookRepository: bookRepository,
	}
}

func (h *bookHandler) AddBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}

	result, err := h.bookRepository.InsertBook(&book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("document with %v id has been added", result)))
}

func (h *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}

	result, err := h.bookRepository.UpdateBook(&book)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v document with %v id updated", result, book.ID)))
}

func (h *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var book domain.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}
	result, err := h.bookRepository.DeleteBook(book.ID)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v document with %v id deleted", result, book.ID)))
}

func (h *bookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("content-type","application/json")
	idParam := r.FormValue("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	book, err := h.bookRepository.GetBookByID(id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func (h *bookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("content-type","application/json")

	books, err := h.bookRepository.GetBooks()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

