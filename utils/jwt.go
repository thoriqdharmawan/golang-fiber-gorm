package utils

import (
	"fmt"
	"golang-fiber-gorm/model/entity"
	"strings"
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

func VerifyJWTToken(tokenString string) (*Claims, error) {
	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC-SHA256 (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key used for verification
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid and contains the claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// Function to verify JWT token and handle "Bearer" prefix
func VerifyJWTTokenHandler(token string) error {
	if token == "" {
		return fmt.Errorf("missing authorization header")
	}

	// Step 2: Ensure token has the "Bearer " prefix
	if !strings.HasPrefix(token, "Bearer ") {
		return fmt.Errorf("invalid authorization header format")
	}

	// Step 3: Extract the token after "Bearer " (7 characters long)
	tokenString := token[7:]

	// Step 4: Verify the token
	_, err := VerifyJWTToken(tokenString)

	if err != nil {
		return fmt.Errorf("invalid token")
	}

	return nil
}
