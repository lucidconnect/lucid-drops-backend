package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang-jwt/jwt"
	"github.com/lucidconnect/inverse/services"
)

type contextKey struct {
	name string
}

type AuthDetails struct {
	Address common.Address
}

type PrivyClaims struct {
	AppId      string `json:"aud,omitempty"`
	Expiration uint64 `json:"exp,omitempty"`
	Issuer     string `json:"iss,omitempty"`
	UserId     string `json:"sub,omitempty"`
}

func (c *PrivyClaims) Valid() error {
	if c.AppId != os.Getenv("PRIVY_APP_ID") {
		return fmt.Errorf("aud claim must be your Privy App ID its currently")
	}

	if c.Issuer != "privy.io" {
		return errors.New("iss claim must be 'privy.io'")
	}
	// 🤡 TODO add token expires after testing
	// if c.Expiration < uint64(time.Now().Unix()) {
	// 	return errors.New("token is expired")
	// }

	return nil
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != "ES256" {
		return nil, fmt.Errorf("unexpected JWT signing method=%v", token.Header["alg"])
	}

	// https://pkg.go.dev/github.com/dgrijalva/jwt-go#ParseECPublicKeyFromPEM
	verificationKey := fmt.Sprint(os.Getenv("PRIVY_VERIFICATION_KEY"))
	// 	verificationKey := `-----BEGIN PUBLIC KEY-----
	// MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEHNoGaXIavvTyGZjULmXXD2TZnxnR5qJI/U0vodti4LmOdX9kNg1+ioQp1MGtEJUww/FD10RaV+NFqCtp77kHHw==
	// -----END PUBLIC KEY-----`

	pk, err := jwt.ParseECPublicKeyFromPEM([]byte(verificationKey))
	if err != nil {
		return err, nil
	}

	return pk, nil
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

				jwtParts := strings.Split(contextMap["authHeader"].(string), ".")
				if len(jwtParts) != 3 {
					next.ServeHTTP(w, r)
					return
				}

				type MiniInfo struct {
					Iss string `json:"iss"`
				}
				metadata := &MiniInfo{}

				rawDecodedText, _ := base64.RawStdEncoding.DecodeString(jwtParts[1])
				err := json.Unmarshal([]byte(rawDecodedText), metadata)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				if metadata.Iss == "https://api-auth.web3auth.io" {
					contextMap["provider"] = "web3Auth"
				} else {
					contextMap["provider"] = "privy"
				}

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
	case "privy":
		var c PrivyClaims

		token, err := jwt.ParseWithClaims(claims["authHeader"].(string), &c, keyFunc)
		if err != nil {
			return nil, err
		}

		// Parse the JWT claims into our custom struct
		privyClaim, ok := token.Claims.(*PrivyClaims)
		if !ok {
			fmt.Println("JWT does not have all the necessary claims.")
		}

		// Check the JWT claims
		err = c.Valid()
		if err != nil {
			return nil, err
		}

		privyWallet, err := GetPrivyWalletsFromSubKey(privyClaim.UserId)
		if err != nil {
			return nil, err
		}

		info.Address = *privyWallet

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
			if jwtInfo.Wallets[0].Address == "" {
				// verify that there's a way to get the address from the public key
				// deriv = jwtInfo.Wallets[0].PublicKey
				isExt = false
				info.Address = common.HexToAddress(fmt.Sprintf("0x%s", jwtInfo.Wallets[0].PublicKey))
			} else {
				info.Address = common.HexToAddress(jwtInfo.Wallets[0].Address)
			}

			// TODO add JWT verification and assert address is present before proceeding
			_, err = services.VerifyWeb3AuthKey(claims["authHeader"].(string), isExt)
			if err != nil {
				return nil, err
			}
		} else {
			parsedAddres := common.HexToAddress(interalJWT.Address)
			info.Address = parsedAddres
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

		parsedAddress := common.HexToAddress(*publicAddress)
		info.Address = parsedAddress
	}

	return &info, nil
}
