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
}

func NewUserService(userRepository entity.UserRepository, validate *validator.Validate) entity.UserService {
	return &UserService{UserRepository: userRepository, Validate: validate}
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

	if err = userService.UserRepository.Register(user); err != nil {
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

func (userService *UserService) Login(payload request.UserLogin) (response.UserLogin, error) {
	panic("not implemented") // TODO: Implement
}

func (userService *UserService) Update(id string, payload request.UserUpdate) (response.UserUpdate, error) {
	panic("not implemented") // TODO: Implement
}

func (userService *UserService) Delete(id string) (response.UserDelete, error) {
	panic("not implemented") // TODO: Implement
}
