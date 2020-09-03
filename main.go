package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lucasasoaresmar/golang-pos/features/auth"
	authMid "github.com/lucasasoaresmar/golang-pos/features/auth/middlewares"
	"github.com/lucasasoaresmar/golang-pos/features/pos"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/ping", authMid.Auth(ping))

	v1Router := router.PathPrefix("/api").Subrouter()
	v1Router.Use(defaultHeaderMiddleware)

	auth.RegisterRoutes(v1Router)
	pos.RegisterRoutes(v1Router)

	log.Fatal(http.ListenAndServe(port(), router))
}

func ping(w http.ResponseWriter, req *http.Request) {
	user, _ := authMid.ContextUser(req)
	fmt.Println(user.Name, user.ID, user.Email, user.Password, user.Roles)
	fmt.Fprintf(w, "pong")
}

func port() string {
	env, ok := os.LookupEnv("PORT")
	if ok {
		return env
	}
	return ":8000"
}
