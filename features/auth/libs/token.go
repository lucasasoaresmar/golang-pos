package libs

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/lucasasoaresmar/golang-pos/features/auth/models"
)

// CustomClaims type ...
type CustomClaims struct {
	*models.User
	jwt.StandardClaims
}

// TokenService library
type TokenService struct{}

// Create a token
func (ts *TokenService) Create(user *models.User) (tokenString string, err error) {
	claims := CustomClaims{
		user,
		jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(tokenSecretKey())

	return tokenString, err
}

// Parse a token
func (ts *TokenService) Parse(tokenString string) (claims interface{}, err error) {
	parsedToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, claimsValidation)
	if err != nil {
		return
	}

	claims, ok := parsedToken.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("There was an error while parsing your token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("Invalid Token")
	}

	return claims, nil
}

func tokenSecretKey() []byte {
	env, ok := os.LookupEnv("TOKEN_SECRET_KEY")
	if ok {
		return []byte(env)
	}
	return []byte("SomeSecreteKey")
}

func claimsValidation(tk *jwt.Token) (interface{}, error) {
	return tokenSecretKey(), nil
}
