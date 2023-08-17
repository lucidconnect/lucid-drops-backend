package internal

import (
	"time"

	"github.com/golang-jwt/jwt"
	"inverse.so/utils"
)

func GenerateJWT(address string) (string, error) {
	jwtSignature := utils.UseEnvOrDefault("JWT_KEY", "0xdd05894fa54754e5ca1add04fe3e970c1171126c6d93f579a5c76cb3f658c3b6")

	currentTime := time.Now()

	ttl := currentTime.Add(6 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     ttl,
		"address": address,
	})

	tokenString, err := token.SignedString([]byte(jwtSignature))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
