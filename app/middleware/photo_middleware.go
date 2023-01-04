package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/response"
)

func PhotoMiddleware(photoService entity.PhotoService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var (
				photo response.PhotoGetOne
				err   error
			)

			ctxKeyUser := helper.ContextKey("user")
			user := r.Context().Value(ctxKeyUser).(*helper.UserClaims)
			userID := user.UserID
			stringPhotoID := chi.URLParam(r, "photoID")
			photoID := uuid.MustParse(stringPhotoID)
			encoder := json.NewEncoder(w)

			w.Header().Add("Content-Type", "application/json")

			if photo, err = photoService.GetOne(photoID); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				encoder.Encode(response.Failure{
					Errors: response.Message{
						Message: exception.ErrEntityNotFound.Error(),
					},
				})

				return
			}

			if photo.UserID != userID {
				w.WriteHeader(http.StatusInternalServerError)
				encoder.Encode(response.Failure{
					Errors: response.Message{
						Message: exception.ErrUnauthorized.Error(),
					},
				})

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
