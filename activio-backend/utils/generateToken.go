package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JWT_EXPIRATION = 24 * time.Hour
const JWT_SECRET = "secret" // TODO: Change this to an environment variable

func GenerateToken(id uint) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(JWT_EXPIRATION).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
