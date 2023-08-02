package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/MicahParks/keyfunc"
	"github.com/golang-jwt/jwt/v4"
	"go.step.sm/crypto/jose"
)

func VerifyWeb3AuthKey(authHeader string, isExternalWallet bool) (bool, error) {

	jwksURL := "https://api.openlogin.com/jwks"
	if isExternalWallet {
		jwksURL = "https://authjs.web3auth.io/jwks"
	}

	ctx, cancel := context.WithCancel(context.Background())
	options := keyfunc.Options{
		Ctx: ctx,
		RefreshErrorHandler: func(err error) {
			log.Printf("There was an error with the jwt.Keyfunc\nError: %s", err.Error())
		},
		RefreshInterval:   time.Hour,
		RefreshRateLimit:  time.Minute * 5,
		RefreshTimeout:    time.Second * 10,
		RefreshUnknownKID: true,
	}

	jwks, err := keyfunc.Get(jwksURL, options)
	if err != nil {
		cancel()
		return false, err
	}

	token, err := jwt.Parse(authHeader, jwks.Keyfunc)
	if err != nil {
		cancel()
		return false, err
	}

	if !token.Valid {
		cancel()
		return false, errors.New("invalid token")
	}

	cancel()
	return true, nil
}

func VerifyWeb3AuthKeyCustom(authHeader string) (interface{}, error) {

	jwtBody, err := jose.ParseSigned(authHeader)
	if err != nil {
		return nil, err
	}

	err = jose.Verify(jwtBody, nil, nil)
	if err != nil {
		return nil, err
	}

	return jwtBody, nil
}

type Keys struct {
	Keys []struct {
		Kty string `json:"kty"`
		Crv string `json:"crv"`
		X   string `json:"x"`
		Y   string `json:"y"`
		Kid string `json:"kid"`
	} `json:"keys"`
}
