package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/domain"
	"server/infrastructure"
	"strconv"
)

type ReaderHandler interface {
	AddReader(w http.ResponseWriter, r *http.Request)
	UpdateReader(w http.ResponseWriter, r *http.Request)
	DeleteReader(w http.ResponseWriter, r *http.Request)
	GetReaderByID(w http.ResponseWriter, r *http.Request)
	GetReaders(w http.ResponseWriter, r *http.Request)
}

type readerHandler struct {
	readerRepository infrastructure.ReaderRepository
}

func NewReaderHandler(readerRepository infrastructure.ReaderRepository) ReaderHandler {
	return &readerHandler{
		readerRepository: readerRepository,
	}
}

func (h *readerHandler) AddReader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var reader domain.Reader
	if err := json.NewDecoder(r.Body).Decode(&reader); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}

	result, err := h.readerRepository.InsertReader(&reader)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("document with %v id has been added", result)))
}

func (h *readerHandler) UpdateReader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var reader domain.Reader
	if err := json.NewDecoder(r.Body).Decode(&reader); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}

	result, err := h.readerRepository.UpdateReader(&reader)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v document with %v id updated", result, reader.ID)))
}

func (h *readerHandler) DeleteReader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var reader domain.Reader
	if err := json.NewDecoder(r.Body).Decode(&reader); err != nil {
		http.Error(w, "Error during parsing body", http.StatusBadRequest)
		return
	}
	result, err := h.readerRepository.DeleteReader(reader.ID)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%v document with %v id deleted", result, reader.ID)))
}

func (h *readerHandler) GetReaderByID(w http.ResponseWriter, r *http.Request) {
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

	reader, err := h.readerRepository.GetReaderByID(id)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(reader)
}

func (h *readerHandler) GetReaders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Add("content-type","application/json")

	readers, err := h.readerRepository.GetReaders()
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(readers)
}

