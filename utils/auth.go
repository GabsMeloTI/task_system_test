package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// HashPassword hashes a given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain password to check if they match.
func ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

var jwtSecretKey = []byte("your_secret_key")

// JWTClaims represents the claims in the JWT token.
type JWTClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// IsAuthenticated checks if the user is authenticated based on the JWT token.
func IsAuthenticated(r *http.Request) (*JWTClaims, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, errors.New("authentication token not found")
	}

	tokenStr := cookie.Value
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
