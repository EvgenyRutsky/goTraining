package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"httpserver/domain"
	"httpserver/infrastructure/client"
	"testing"
)

func TestNewBookRepository_ShouldNotReturnNil(t *testing.T) {
	cfg := client.NewConfig()
	b := NewBookRepository(cfg)
	assert.NotNil(t, b)
}

func TestBooksRepository_InsertBookShouldReturnBookIDAndNilError(t *testing.T) {

	testBook := &domain.Book{
		ID:     123,
		Name:   "testBook",
		Author: "testAuthor",
	}
	cfg := client.NewConfig()
	br := NewBookRepository(cfg)
	result, err := br.InsertBook(testBook)
	assert.Equal(t, 123, result)
	assert.Nil(t, err)
}

func TestBookRepository_UpdateBookShouldReturnBookIDAndNilError(t *testing.T) {
	testBook := &domain.Book{
		ID:     123,
		Name:   "testBook2",
		Author: "testAuthor2",
	}
	cfg := client.NewConfig()
	br := NewBookRepository(cfg)
	result, err := br.UpdateBook(testBook)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestBooksRepository_GetBookByIDShouldReturnExistingBook(t *testing.T) {
	testBookID := 123
	cfg := client.NewConfig()
	br := NewBookRepository(cfg)
	result, err := br.GetBookByID(testBookID)
	assert.Equal(t, testBookID, result.ID)
	assert.Nil(t, err)
}

func TestBooksRepository_GetBookShouldReturnASliceOfBooks(t *testing.T) {
	cfg := client.NewConfig()
	br := NewBookRepository(cfg)
	result, err := br.GetBooks()
	assert.NotNil(t,result)
	assert.Nil(t, err)
}

func TestBookRepository_DeleteBookShouldReturnDeletedID(t *testing.T) {
	testBookID := 123
	cfg := client.NewConfig()
	br := NewBookRepository(cfg)
	result, err := br.DeleteBook(testBookID)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestDeleteBookWithIncorrectConfig(t *testing.T) {
	testBookID := 123
	cfg := &client.Config{
		ClientURI:         "123",
		Dbname:            "123",
		BooksCollection:   "123",
		ReadersCollection: "123",
	}
	br := NewBookRepository(cfg)
	result, err := br.DeleteBook(testBookID)
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
}