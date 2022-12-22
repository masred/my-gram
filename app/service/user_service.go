package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type UserService struct {
	UserRepository entity.UserRepository
	Validate       *validator.Validate
	UserJWTService *helper.UserJWTService
}

func NewUserService(userRepository entity.UserRepository, validate *validator.Validate, userJWTService *helper.UserJWTService) entity.UserService {
	return &UserService{UserRepository: userRepository, Validate: validate, UserJWTService: userJWTService}
}

func (userService *UserService) Register(req request.UserRegister) (res response.UserRegister, err error) {

	if err = userService.Validate.Struct(req); err != nil {
		return
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return
	}

	user := entity.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Age:      req.Age,
	}

	if err = userService.UserRepository.Register(&user); err != nil {
		return
	}

	res = response.UserRegister{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return
}

func (userService *UserService) Login(req request.UserLogin) (res response.UserLogin, err error) {
	if err = userService.Validate.Struct(req); err != nil {
		return
	}

	user := entity.User{
		Email:    req.Email,
		Password: []byte(req.Password),
	}

	if err = userService.UserRepository.Login(&user); err != nil {
		return
	}

	tokenResponse, err := userService.UserJWTService.GenerateUserToken(user.ID, user.Username, user.Email)

	res = response.UserLogin{
		Token: tokenResponse,
	}

	return
}

func (userService *UserService) Update(id uuid.UUID, req request.UserUpdate) (res response.UserUpdate, err error) {
	if err = userService.Validate.Struct(req); err != nil {
		return
	}

	user := entity.User{
		ID:       id,
		Username: req.Username,
		Email:    req.Email,
	}

	if err = userService.UserRepository.Update(&user); err != nil {
		return
	}

	res = response.UserUpdate{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Age:       user.Age,
		UpdatedAt: user.UpdatedAt,
	}

	return
}

func (userService *UserService) Delete(id uuid.UUID) (res response.UserDelete, err error) {
	if err = userService.UserRepository.Delete(id); err != nil {
		return
	}

	res = response.UserDelete{
		Message: "user successfully deleted",
	}

	return
}
