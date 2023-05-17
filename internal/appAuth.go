package internal

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type contextKey struct {
	name string
}

type AuthDetails struct {
	Address string
}

var (
	userAuthToken       = &contextKey{"userAuthToken"}
	errJWTCreationError = errors.New("authentication Failed")
)

func UserAuthMiddleWare() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// secretKey := utils.UseEnvOrDefault("JWT_SECRET_KEY", "n0t50r4n60m")
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

			jwtToken := authHeader

			jwtParts := strings.Split(jwtToken, ".")

			// if len(jwtParts) != 3 {
			// 	// 	// All authentication related checks will be done within the various queries and mutations
			// 	// 	// w.WriteHeader(http.StatusUnauthorized)
			// 	// 	// _ = json.NewEncoder(w).Encode(map[string]interface{}{
			// 	// 	// 	"message": "Unauthorized",
			// 	// 	// })
			// 	// 	next.ServeHTTP(w, r)
			// 	// 	return
			// }

			rawDecodedText, _ := base64.URLEncoding.DecodeString(jwtParts[1])
			if string(rawDecodedText) != "" {
				ctx := context.WithValue(r.Context(), userAuthToken, rawDecodedText)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
			}
		})
	}
}

func GetAuthDetailsFromContext(ctx context.Context) (authDetails *DynamicJWTMetadata, err error) {
	jwtClaims, ok := ctx.Value(userAuthToken).([]byte)
	if !ok {
		return nil, errJWTCreationError
	}

	var jwtInfo DynamicJWTMetadata
	err = json.Unmarshal(jwtClaims, &jwtInfo)
	if err != nil {
		return nil, err
	}
	// jwtInfo, casted := claims["user_id"].(DynamicJWTMetadata)
	// if !casted {
	// 	return nil, errJWTCreationError
	// }

	return &jwtInfo, nil
}
