package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"httpserver/domain"
	"httpserver/infrastructure"
	"httpserver/infrastructure/client"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewBookHandler(t *testing.T) {
	cfg := client.NewConfig()
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	assert.NotNil(t, h)
}

var testbooksdb = "test"

func TestAddBook(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook3",
		Author: "testAuthor3",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testBook)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.AddBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/books/addbook", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("document with %v id has been added", testBook.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestAddBookWithIncorrectMethod(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook3",
		Author: "testAuthor3",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testBook)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.AddBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet,"/books/addbook", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusMethodNotAllowed, rec.Code)
}

func TestAddBookWithIncorrectBody(t *testing.T) {
	testBook := "IncorrectBody"
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testBook)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.AddBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/books/addbook", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestAddBookWithIncorrectDBConfig(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook3",
		Author: "testAuthor3",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testBook)
	cfg := &client.Config{
		ClientURI:         "123123",
		Dbname:            "123123",
		BooksCollection:   "123123",
		ReadersCollection: "123123",
	}
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.AddBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/books/addbook", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestUpdateBook(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook34",
		Author: "testAuthor34",
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testBook)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.UpdateBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/books/updatebook", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("%v document with %v id updated", 1, testBook.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBookByID(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook34",
		Author: "testAuthor34",
	}
	expected := new(bytes.Buffer)
	json.NewEncoder(expected).Encode(testBook)
	testBookID := testBook.ID
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.GetBookByID)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/books/getbook?id=%v", testBookID), nil)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, expected.Bytes(), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBooks(t *testing.T) {
	testBooks := []domain.Book{
		{
			ID:     1234,
			Name:   "testBook34",
			Author: "testAuthor34",
		},
	}
	expected := new(bytes.Buffer)
	json.NewEncoder(expected).Encode(testBooks)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.GetBooks)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/books/getbooks", nil)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, expected.Bytes(), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteBook(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook34",
		Author: "testAuthor34",
	}
	testBody := new(bytes.Buffer)
	json.NewEncoder(testBody).Encode(testBook)
	cfg := client.NewConfig()
	cfg.Dbname = testbooksdb
	br := infrastructure.NewBookRepository(cfg)
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.DeleteBook)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/books/deletebook", testBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("%v document with %v id deleted", 1, testBook.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}
