package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

type transactionRepository interface {
	GetAll() (categories []models.Transaction, err error)
	Create(transaction *models.Transaction) error
	Update(transaction *models.Transaction) error
	Delete(id string) error
}

// TransactionControllers handles http requests
type TransactionControllers struct {
	TransactionRepository transactionRepository
}

// GetAll product categories
func (tctrl *TransactionControllers) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := tctrl.TransactionRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Create product Transaction
func (tctrl *TransactionControllers) Create(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = tctrl.TransactionRepository.Create(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// Update product Transaction
func (tctrl *TransactionControllers) Update(w http.ResponseWriter, r *http.Request) {
	transaction := models.Transaction{}
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = tctrl.TransactionRepository.Update(&transaction)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transaction)
}

// Delete product Transaction
func (tctrl *TransactionControllers) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := tctrl.TransactionRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
