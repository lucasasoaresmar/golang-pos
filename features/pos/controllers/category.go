package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasasoaresmar/golang-pos/features/pos/models"
)

type categoryRepository interface {
	GetAll() (categories []models.Category, err error)
	Create(category *models.Category) error
	Update(category *models.Category) error
	Delete(id string) error
}

// CategoryControllers handles http requests
type CategoryControllers struct {
	CategoryRepository categoryRepository
}

// GetAll product categories
func (cctrl *CategoryControllers) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := cctrl.CategoryRepository.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Create product category
func (cctrl *CategoryControllers) Create(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = cctrl.CategoryRepository.Create(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}

// Update product category
func (cctrl *CategoryControllers) Update(w http.ResponseWriter, r *http.Request) {
	category := models.Category{}
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = cctrl.CategoryRepository.Update(&category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(category)
}

// Delete product category
func (cctrl *CategoryControllers) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := cctrl.CategoryRepository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(nil)
}
