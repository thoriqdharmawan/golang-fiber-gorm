package utils

import (
	"golang-fiber-gorm/model/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = []byte("secret_key")

type Claims struct {
	User entity.User
	jwt.RegisteredClaims
}

func GenerateJWTToken(user entity.User) (string, error) {
	// Define token expiration time (e.g., 1 hour)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the claims with the user data and expiration time
	claims := &Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create a new token with the specified claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
