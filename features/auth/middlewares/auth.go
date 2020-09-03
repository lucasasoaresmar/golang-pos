package middlewares

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/lucasasoaresmar/golang-pos/features/auth/libs"
	"github.com/lucasasoaresmar/golang-pos/features/auth/models"
)

// AuthContextKey to be used in http request context
type AuthContextKey string

// Auth middleware
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := getTokenFromAuthorization(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		tokenService := libs.TokenService{}
		claims, err := tokenService.Parse(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), AuthContextKey("auth"), claims)
		next(w, r.WithContext(ctx))
	}
}

// OnlyAdmins check if token belongs to an admin
func OnlyAdmins(next http.HandlerFunc) http.HandlerFunc {
	return Auth(func(w http.ResponseWriter, r *http.Request) {
		user, err := ContextUser(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if !user.IsValidRole("admin") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Only administrators have permission to do this"))
			return
		}

		next(w, r)
	})
}

// ContextUser ...
func ContextUser(req *http.Request) (user models.User, err error) {
	authContext := req.Context().Value(AuthContextKey("auth"))
	if authContext != nil {
		claims := authContext.(*libs.CustomClaims)
		user = *claims.User
		return
	}

	err = errors.New("There is no authentication context in this request")
	return
}

func getTokenFromAuthorization(r *http.Request) (tokenString string, err error) {
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(authHeader) != 2 {
		err = errors.New("Malformed Authorization header")
		return
	}

	tokenString = authHeader[1]
	return
}
