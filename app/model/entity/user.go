package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type (
	User struct {
		ID          uuid.UUID     `gorm:"primaryKey"`
		Username    string        `gorm:"not null;uniqueIndex"`
		Email       string        `gorm:"not null;uniqueIndex"`
		Password    []byte        `gorm:"not null"`
		Age         int           `gorm:"not null"`
		Photos      []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	UserService interface {
		Register(payload request.UserRegister) (response.UserRegister, error)
		Login(payload request.UserLogin) (response.UserLogin, error)
		Update(id uuid.UUID, payload request.UserUpdate) (response.UserUpdate, error)
		Delete(id uuid.UUID) (response.UserDelete, error)
	}

	UserRepository interface {
		Register(user *User) error
		Login(user *User) error
		Update(id uuid.UUID, payload request.UserUpdate) error
		Delete(id uuid.UUID) error
	}
)
