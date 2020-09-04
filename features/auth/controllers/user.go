package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasasoaresmar/golang-pos/features/auth/models"
)

type userRepository interface {
	Create(user *models.User) error
	GetByEmailAndPassword(email string, password string) (models.User, error)
}

type tokenService interface {
	Create(user *models.User) (token string, err error)
}

// UserControllers handles http requests
type UserControllers struct {
	UserRepository userRepository
	TokenService   tokenService
}

// Create a new User
func (uc *UserControllers) Create(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	err = uc.UserRepository.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Login an user
func (uc *UserControllers) Login(w http.ResponseWriter, r *http.Request) {
	var request = struct {
		Email    string
		Password string
	}{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	if request.Email == "" || request.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Must supply email and password"))
		return
	}

	user, err := uc.UserRepository.GetByEmailAndPassword(request.Email, request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tokenString, err := uc.TokenService.Create(&user)

	response, err := json.Marshal(map[string]string{"token": tokenString})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(response))
}
