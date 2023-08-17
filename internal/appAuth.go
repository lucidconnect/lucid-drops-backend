package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"inverse.so/services"
)

type contextKey struct {
	name string
}

type AuthDetails struct {
	Address string
}

var (
	userAuthToken = &contextKey{"userAuthToken"}
)

func UserAuthMiddleWare() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				next.ServeHTTP(w, r)
				return
			}

			if strings.HasPrefix(authHeader, "Bearer ") {
				authHeader = authHeader[7:]
			} else {
				next.ServeHTTP(w, r)
				return
			}

			var contextMap = make(map[string]interface{})
			switch authHeader[:1] {
			case "e":
				jwtToken := authHeader

				contextMap["authHeader"] = jwtToken
				contextMap["provider"] = "web3Auth"
				ctx := context.WithValue(r.Context(), userAuthToken, contextMap)
				next.ServeHTTP(w, r.WithContext(ctx))

			case "W":
				contextMap["authHeader"] = authHeader
				contextMap["provider"] = "magic"
				ctx := context.WithValue(r.Context(), userAuthToken, contextMap)
				next.ServeHTTP(w, r.WithContext(ctx))
			}
		})
	}
}

func GetAuthDetailsFromContext(ctx context.Context) (authDetails *AuthDetails, err error) {

	claims, ok := ctx.Value(userAuthToken).(map[string]interface{})
	if !ok {
		return nil, errors.New("jwt claims not found in context")
	}

	provider, ok := claims["provider"].(string)
	if !ok {
		return nil, errors.New("jwt claims not found in context")
	}

	var info AuthDetails
	switch provider {
	case "web3Auth":
		jwtParts := strings.Split(claims["authHeader"].(string), ".")
		rawDecodedText, _ := base64.RawStdEncoding.DecodeString(jwtParts[1])

		var interalJWT CustomJWTMetadata
		err = json.Unmarshal(rawDecodedText, &interalJWT)
		if interalJWT.Address == "" || err != nil {
			var jwtInfo Web3AuthMetadata
			err = json.Unmarshal(rawDecodedText, &jwtInfo)
			if err != nil {
				return nil, err
			}

			isExt := true
			info.Address = jwtInfo.Wallets[0].Address
			if info.Address == "" {
				// verify that there's a way to get the address from the public key
				// deriv = jwtInfo.Wallets[0].PublicKey
				isExt = false
				info.Address = fmt.Sprintf("0x%s", jwtInfo.Wallets[0].PublicKey)
			}

			// TODO add JWT verification and assert address is present before proceeding
			_, err = services.VerifyWeb3AuthKey(claims["authHeader"].(string), isExt)
			if err != nil {
				return nil, err
			}
		} else {
			info.Address = interalJWT.Address
		}

	case "magic":
		jwtClaims, ok := claims["authHeader"].(string)
		if !ok {
			return nil, errors.New("jwt claims not found in context")
		}
		magicPayload, err := services.GenerateMagicJWT(string(jwtClaims))
		if err != nil {
			return nil, err
		}

		publicAddress, err := services.GetMagicAddress(magicPayload)
		if err != nil {
			return nil, err
		}

		info.Address = *publicAddress
	}

	return &info, nil
}
