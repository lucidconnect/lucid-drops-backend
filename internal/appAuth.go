package internal

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

type contextKey struct {
	name string
}

type AuthDetails struct {
	DeviceId string
	UserUUID string
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

			rawDecodedText, err := base64.StdEncoding.DecodeString(jwtParts[0])
			if err == nil {
				ctx := context.WithValue(r.Context(), userAuthToken, rawDecodedText)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
			}
		})
	}
}

func GetAuthDetailsFromContext(ctx context.Context) (authDetails *AuthDetails, err error) {
	claims, ok := ctx.Value(userAuthToken).(jwt.MapClaims)
	if !ok {
		return nil, errJWTCreationError
	}

	userUUID, casted := claims["user_id"].(string)
	if !casted {
		return nil, errJWTCreationError
	}

	deviceId, casted := claims["device_id"].(string)
	if !casted {
		return nil, errJWTCreationError
	}

	return &AuthDetails{
		UserUUID: userUUID,
		DeviceId: deviceId,
	}, nil
}
