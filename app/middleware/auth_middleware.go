package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/model/http/response"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")
		encoder := json.NewEncoder(w)

		w.Header().Add("Content-Type", "application/json")

		if !bearer {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(response.Failure{
				Errors: response.Message{
					Message: exception.ErrUnauthorized.Error(),
				},
			})

			return
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")
		if secretKey == "" {
			secretKey = "ez"
		}

		jwtService := helper.NewUserJWTService([]byte(secretKey))

		stringToken := strings.Split(headerToken, " ")[1]

		parsedToken, err := jwtService.ParseUserToken(stringToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			encoder.Encode(response.Failure{
				Errors: response.Message{
					Message: exception.ErrUnauthorized.Error(),
				},
			})

			return
		}

		ctxKeyUser := helper.ContextKey("user")
		ctx := context.WithValue(r.Context(), ctxKeyUser, parsedToken.Claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
