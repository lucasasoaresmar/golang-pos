package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

type itemRepository interface {
	GetAll() (categories []models.Item, err error)
	Create(item *models.Item) error
	Update(item *models.Item) error
	Delete(id string) error
}

// ItemControllers handles http requests
type ItemControllers struct {
	ItemRepository itemRepository
}

// GetAll product categories
func (ictrl *ItemControllers) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := ictrl.ItemRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Create product Item
func (ictrl *ItemControllers) Create(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = ictrl.ItemRepository.Create(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// Update product Item
func (ictrl *ItemControllers) Update(w http.ResponseWriter, r *http.Request) {
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = ictrl.ItemRepository.Update(&item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// Delete product Item
func (ictrl *ItemControllers) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := ictrl.ItemRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
