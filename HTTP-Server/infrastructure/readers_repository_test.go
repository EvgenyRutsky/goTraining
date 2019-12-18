package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"httpserver/domain"
	"httpserver/infrastructure/client"
	"testing"
)

func TestNewReadersRepository_ShouldNotReturnNil(t *testing.T) {
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	assert.NotNil(t, r)
}

func TestReadersRepository_InsertReaderShouldReturnReaderIDAndNilError(t *testing.T) {

	testReader := &domain.Reader{
		ID:     123,
		Name:   "testReader",
		BookID: 321,
	}
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	result, err := r.InsertReader(testReader)
	assert.Equal(t, 123, result)
	assert.Nil(t, err)
}

func TestReadersRepository_UpdateReaderShouldReturnReaderIDAndNilError(t *testing.T) {
	testReader := &domain.Reader{
		ID:     123,
		Name:   "testReader1",
		BookID: 321,
	}
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	result, err := r.UpdateReader(testReader)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestReadersRepository_GetReaderByIDShouldReturnExistingReader(t *testing.T) {
	testReaderID := 123
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	result, err := r.GetReaderByID(testReaderID)
	assert.Equal(t, testReaderID, result.ID)
	assert.Nil(t, err)
}

func TestReadersRepository_GetReadersShouldReturnASliceOfReaders(t *testing.T) {
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	result, err := r.GetReaders()
	assert.NotNil(t,result)
	assert.Nil(t, err)
}

func TestReadersRepository_DeleteReaderShouldReturnDeletedID(t *testing.T) {
	testReaderID := 123
	cfg := client.NewConfig()
	r := NewReaderRepository(cfg)
	result, err := r.DeleteReader(testReaderID)
	assert.Equal(t, 1, result)
	assert.Nil(t, err)
}

func TestDeleteReaderWithIncorrectConfig(t *testing.T) {
	testReaderID := 123
	cfg := &client.Config{
		ClientURI:         "123",
		Dbname:            "123",
		BooksCollection:   "123",
		ReadersCollection: "123",
	}
	r := NewReaderRepository(cfg)
	result, err := r.DeleteReader(testReaderID)
	assert.Equal(t, 0, result)
	assert.NotNil(t, err)
}
