package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"server/domain"
	"testing"
)

func TestNewReadersRepository_ShouldNotReturnNil(t *testing.T) {
	r := NewReaderRepository()
	assert.NotNil(t, r)
}

func TestReadersRepository_InsertReaderShouldReturnReaderIDAndNilError(t *testing.T) {

	testReader := &domain.Reader{
		ID:     123,
		Name:   "testReader",
		BookID: 321,
	}

	r := NewReaderRepository()
	result, err := r.InsertReader(testReader)
	assert.Equal(t, 123, result)
	assert.Nil(t, err)
}

func TestNewReadersRepository_UpdateReaderShouldReturnReaderIDAndNilError(t *testing.T) {
	testReader := &domain.Reader{
		ID:     123,
		Name:   "testReader",
		BookID: 321,
	}

	r := NewReaderRepository()
	result, err := r.UpdateReader(testReader)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestReadersRepository_GetReaderByIDShouldReturnExistingReader(t *testing.T) {
	testReaderID := 123
	r := NewReaderRepository()
	result, err := r.GetReaderByID(testReaderID)
	assert.Equal(t, testReaderID, result.ID)
	assert.Nil(t, err)
}

func TestReadersRepository_GetReadersShouldReturnASliceOfReaders(t *testing.T) {
	r := NewReaderRepository()
	result, err := r.GetReaders()
	assert.NotNil(t,result)
	assert.Nil(t, err)
}

func TestReadersRepository_DeleteReaderShouldReturnDeletedID(t *testing.T) {
	testReaderID := 123
	r := NewReaderRepository()
	result, err := r.DeleteReader(testReaderID)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}
