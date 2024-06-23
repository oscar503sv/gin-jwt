package utils

import (
	"time"

	"github.com/oscar503sv/gin-jwt/services"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("secret")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(12 * time.Hour)
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(signedToken string) (bool, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return false, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return false, err
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return false, err
	}
	return true, nil
}

func RemoveToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) {
		return err
	}

	// Añadir el token a la lista negra
	if err := services.AddToBlacklist(signedToken); err != nil {
		return err
	}
	return nil
}
