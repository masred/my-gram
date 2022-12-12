package repository

import (
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
func (userRepository *UserRepository) Register(user entity.User) (err error) {
	if err = userRepository.Database.Create(&user).Error; err != nil {
		return
	}

	return
}

func (userRepository *UserRepository) Login(user entity.User) error {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Update(id string, payload request.UserUpdate) error {
	panic("not implemented") // TODO: Implement
}

func (userRepository *UserRepository) Delete(id string) error {
	panic("not implemented") // TODO: Implement
}
