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

func TestNewReaderHandler(t *testing.T) {
	cfg := client.NewConfig()
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	assert.NotNil(t, h)
}

var testreadersdb = "test"

func TestAddReader(t *testing.T) {
	testReader := domain.Reader{
		ID:     1234,
		Name:   "testReader3",
		BookID: 12345,
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testReader)
	cfg := client.NewConfig()
	cfg.Dbname = testreadersdb
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	handler := http.HandlerFunc(h.AddReader)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/readers/addreader", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("document with %v id has been added", testReader.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestUpdateReader(t *testing.T) {
	testReader := domain.Reader{
		ID:     1234,
		Name:   "testReader34",
		BookID: 12345,
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(testReader)
	cfg := client.NewConfig()
	cfg.Dbname = testreadersdb
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	handler := http.HandlerFunc(h.UpdateReader)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost,"/readers/updatereader", reqBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("%v document with %v id updated", 1, testReader.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetReaderByID(t *testing.T) {
	testReader := domain.Reader{
		ID:     1234,
		Name:   "testReader34",
		BookID: 12345,
	}
	expected := new(bytes.Buffer)
	json.NewEncoder(expected).Encode(testReader)
	cfg := client.NewConfig()
	cfg.Dbname = testreadersdb
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	handler := http.HandlerFunc(h.GetReaderByID)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/readers/getreader?id=%v", testReader.ID), nil)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, expected.Bytes(), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestGetReaders(t *testing.T) {
	testReaders := []domain.Reader{
		{
			ID:     1234,
			Name:   "testReader34",
			BookID: 12345,
		},
	}
	expected := new(bytes.Buffer)
	json.NewEncoder(expected).Encode(testReaders)
	cfg := client.NewConfig()
	cfg.Dbname = testreadersdb
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	handler := http.HandlerFunc(h.GetReaders)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/readers/getreaders", nil)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, expected.Bytes(), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestDeleteReader(t *testing.T) {
	testReader := domain.Reader{
		ID:     1234,
		Name:   "testReader34",
		BookID: 12345,
	}
	testBody := new(bytes.Buffer)
	json.NewEncoder(testBody).Encode(testReader)
	cfg := client.NewConfig()
	cfg.Dbname = testreadersdb
	r := infrastructure.NewReaderRepository(cfg)
	h := NewReaderHandler(r)
	handler := http.HandlerFunc(h.DeleteReader)
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/readers/deletereader", testBody)
	handler.ServeHTTP(rec, req)
	assert.Equal(t, []byte(fmt.Sprintf("%v document with %v id deleted", 1, testReader.ID)), rec.Body.Bytes())
	assert.Equal(t, http.StatusOK, rec.Code)
}
