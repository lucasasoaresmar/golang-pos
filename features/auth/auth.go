package auth

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasasoaresmar/golang-pos/features/auth/controllers"
	"github.com/lucasasoaresmar/golang-pos/features/auth/libs"
	"github.com/lucasasoaresmar/golang-pos/features/auth/middlewares"
	"github.com/lucasasoaresmar/golang-pos/features/auth/repositories"
)

// RegisterRoutes ...
func RegisterRoutes(router *mux.Router) {
	authRouter := router.PathPrefix("/v1/auth").Subrouter()

	tokenService := libs.TokenService{}

	userRepository := repositories.UserRepository{}
	userHandlers := controllers.UserControllers{
		UserRepository: &userRepository,
		TokenService:   &tokenService,
	}

	authRouter.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	authRouter.HandleFunc("/user", middlewares.OnlyAdmins(userHandlers.Create)).Methods("POST")
	authRouter.HandleFunc("/login", userHandlers.Login).Methods("POST")
}
