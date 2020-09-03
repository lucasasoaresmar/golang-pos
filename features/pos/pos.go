package pos

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	authMid "github.com/lucasasoaresmar/golang-pos/features/auth/middlewares"
	"github.com/lucasasoaresmar/golang-pos/features/pos/controllers"
	"github.com/lucasasoaresmar/golang-pos/features/pos/repositories"
)

// RegisterRoutes ...
func RegisterRoutes(router *mux.Router) {
	posRouter := router.PathPrefix("/v1/pos").Subrouter()
	posRouter.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	// Categories endpoints
	categoryRepository := repositories.CategoryRepository{}
	categoryHandlers := controllers.CategoryControllers{
		CategoryRepository: &categoryRepository,
	}
	posRouter.HandleFunc("/category", authMid.OnlyAdmins(categoryHandlers.GetAll)).Methods("GET")
	posRouter.HandleFunc("/category", authMid.OnlyAdmins(categoryHandlers.Create)).Methods("POST")
	posRouter.HandleFunc("/category", authMid.OnlyAdmins(categoryHandlers.Update)).Methods("PUT")
	posRouter.HandleFunc("/category/{id}", authMid.OnlyAdmins(categoryHandlers.Delete)).Methods("DELETE")

	// Item endpoints
	itemRepository := repositories.ItemRepository{}
	itemHandlers := controllers.ItemControllers{
		ItemRepository: &itemRepository,
	}
	posRouter.HandleFunc("/item", authMid.OnlyAdmins(itemHandlers.GetAll)).Methods("GET")
	posRouter.HandleFunc("/item", authMid.OnlyAdmins(itemHandlers.Create)).Methods("POST")
	posRouter.HandleFunc("/item", authMid.OnlyAdmins(itemHandlers.Update)).Methods("PUT")
	posRouter.HandleFunc("/item/{id}", authMid.OnlyAdmins(itemHandlers.Delete)).Methods("DELETE")

	// Transaction endpoints
	transactionRepository := repositories.TransactionRepository{}
	transactionHandlers := controllers.TransactionControllers{
		TransactionRepository: &transactionRepository,
	}
	posRouter.HandleFunc("/transaction", authMid.OnlyAdmins(transactionHandlers.GetAll)).Methods("GET")
	posRouter.HandleFunc("/transaction", authMid.OnlyAdmins(transactionHandlers.Create)).Methods("POST")
	posRouter.HandleFunc("/transaction", authMid.OnlyAdmins(transactionHandlers.Update)).Methods("PUT")
	posRouter.HandleFunc("/transaction/{id}", authMid.OnlyAdmins(transactionHandlers.Delete)).Methods("DELETE")
}
