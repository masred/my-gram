package api

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/middleware"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type PhotoController struct {
	PhotoService entity.PhotoService
}

func NewPhotoController(photoService entity.PhotoService) PhotoController {
	return PhotoController{PhotoService: photoService}
}

func (photoController *PhotoController) Route(r chi.Router) {
	r.Route("/photos", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/", photoController.Create)

		r.Group(func(r chi.Router) {
			r.Use(middleware.PhotoMiddleware(photoController.PhotoService))
		})
	})
}

func (photoController *PhotoController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)
	ctxKeyUser := helper.ContextKey("user")
	user := r.Context().Value(ctxKeyUser).(*helper.UserClaims)
	userID := user.UserID

	w.Header().Add("Content-Type", "application/json")

	payload := request.PhotoCreate{
		UserID: userID,
	}

	if err := decoder.Decode(&payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(response.Failure{
			Errors: response.Message{
				Message: err.Error(),
			},
		})

		return
	}

	data, err := photoController.PhotoService.Create(payload)
	if err != nil {
		errorsResponse := make(map[string][]interface{})

		if reflect.TypeOf(err).String() == "validator.ValidationErrors" {
			validationErrors := err.(validator.ValidationErrors)
			for _, error := range validationErrors {
				field := strings.ToLower(error.Field())
				errorsResponse[field] = append(errorsResponse[field], exception.ValidationMessage(error).Error())
			}
		}

		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(response.Failure{
			Errors: errorsResponse,
		})

		return
	}

	res := response.Success{
		Data: data,
	}

	w.WriteHeader(http.StatusCreated)

	if err = encoder.Encode(res); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(response.Failure{
			Errors: response.Message{
				Message: err.Error(),
			},
		})

		return
	}
}
