package internal

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(phone string) (string, error) {
	phoneNumber := phone

	currentTime := time.Now()

	ttl := currentTime.Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":          ttl,
		"phone_number": phoneNumber,
	})
	tokenString, err := token.SignedString([]byte(""))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
