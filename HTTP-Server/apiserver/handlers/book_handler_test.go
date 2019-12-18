package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"server/domain"
	"server/infrastructure"
	"testing"
)

func TestNewBookHandler(t *testing.T) {
	br := infrastructure.NewBookRepository()
	h := NewBookHandler(br)
	assert.NotNil(t, h)
}

func TestAddBook(t *testing.T) {
	testBook := domain.Book{
		ID:     1234,
		Name:   "testBook3",
		Author: "testAuthor3",
	}
	br := infrastructure.NewBookRepository()
	h := NewBookHandler(br)
	handler := http.HandlerFunc(h.AddBook)
	rec := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/books/addbook"))
}
