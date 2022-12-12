package api

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type UserController struct {
	UserService entity.UserService
}

func NewUserController(userService *entity.UserService) UserController {
	return UserController{UserService: *userService}
}

func (userController *UserController) Route(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Post("/register", userController.Register)
	})
}

func (userController *UserController) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	w.Header().Add("Content-Type", "application/json")

	var payload request.UserRegister
	if err := decoder.Decode(&payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(response.Failure{
			Errors: response.Message{
				Message: err.Error(),
			},
		})

		return
	}

	data, err := userController.UserService.Register(payload)
	if err != nil {
		errorsResponse := make(map[string][]interface{})

		if strings.Contains(err.Error(), "idx_users_username") {
			errorsResponse["username"] = append(errorsResponse["username"], exception.ErrUsernameAlreadyRegistered.Error())
		}

		if strings.Contains(err.Error(), "idx_users_email") {
			errorsResponse["email"] = append(errorsResponse["email"], exception.ErrEmailAlreadyRegistered.Error())
		}

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
