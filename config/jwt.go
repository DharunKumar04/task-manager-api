package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = []byte("AN9DaJf5zzH5mV8A5hULndFQm+CB9s8Kui3XIr+6wNw=")

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})

	fmt.Println("Parse error:", err)         // Debug
	fmt.Println("Token valid:", token.Valid) // Debug

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
