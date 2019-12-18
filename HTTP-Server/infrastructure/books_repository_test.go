package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"server/domain"
	"testing"
)

func TestNewBookRepository_ShouldNotReturnNil(t *testing.T) {
	b := NewBookRepository()
	assert.NotNil(t, b)
}

func TestBooksRepository_InsertBookShouldReturnBookIDAndNilError(t *testing.T) {

	testBook := &domain.Book{
		ID:     123,
		Name:   "testBook",
		Author: "testAuthor",
	}

	br := NewBookRepository()
	result, err := br.InsertBook(testBook)
	assert.Equal(t, 123, result)
	assert.Nil(t, err)
}

func TestNewBookRepository_UpdateBookShouldReturnBookIDAndNilError(t *testing.T) {
	testBook := &domain.Book{
		ID:     123,
		Name:   "testBook2",
		Author: "testAuthor2",
	}

	br := NewBookRepository()
	result, err := br.UpdateBook(testBook)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestBooksRepository_GetBookByIDShouldReturnExistingBook(t *testing.T) {
	testBookID := 123
	br := NewBookRepository()
	result, err := br.GetBookByID(testBookID)
	assert.Equal(t, testBookID, result.ID)
	assert.Nil(t, err)
}

func TestBooksRepository_GetBookShouldReturnASliceOfBooks(t *testing.T) {
	br := NewBookRepository()
	result, err := br.GetBooks()
	assert.NotNil(t,result)
	assert.Nil(t, err)
}

func TestBookRepository_DeleteBookShouldReturnDeletedID(t *testing.T) {
	testBookID := 123
	br := NewBookRepository()
	result, err := br.DeleteBook(testBookID)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}