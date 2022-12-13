package repository

import (
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/model/entity"
	"github.com/masred/my-gram/app/model/http/request"
	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func NewUserRepository(database *gorm.DB) entity.UserRepository {
	return &UserRepository{Database: database}
}
func (userRepository *UserRepository) Register(user *entity.User) (err error) {
	if err = userRepository.Database.Create(&user).Error; err != nil {
		return
	}

	return
}

func (userRepository *UserRepository) Login(user *entity.User) (err error) {
	password := user.Password

	err = userRepository.Database.Where("email = ?", user.Email).Take(&user).Error

	isEmailValid := err == nil
	isPasswordValid := helper.ComparePasswordHash([]byte(user.Password), []byte(password))

	if !isEmailValid || !isPasswordValid {
		return exception.ErrInvalidCredentials
	}

	return
}

func (userRepository *UserRepository) Update(id uuid.UUID, payload request.UserUpdate) error {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Delete(id uuid.UUID) error {
	panic("not implemented") // TODO: Implement
}
