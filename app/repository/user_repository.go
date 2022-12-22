package repository

import (
	"github.com/google/uuid"
	"github.com/masred/my-gram/app/exception"
	"github.com/masred/my-gram/app/helper"
	"github.com/masred/my-gram/app/model/entity"
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

func (userRepository *UserRepository) Update(user *entity.User) (err error) {
	newUser := entity.User{
		Username: user.Username,
		Email:    user.Email,
	}

	if err = userRepository.Database.First(&user, user.ID).Error; err != nil {
		return
	}

	if err = userRepository.Database.Model(&user).Updates(newUser).Error; err != nil {
		return
	}

	return
}

func (userRepository *UserRepository) Delete(id uuid.UUID) (err error) {
	var user entity.User

	if err = userRepository.Database.First(&user, id).Error; err != nil {
		return
	}

	if err = userRepository.Database.Delete(&user).Error; err != nil {
		return
	}

	return
}
