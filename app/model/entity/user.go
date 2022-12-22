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
		Register(payload request.UserRegister) (res response.UserRegister, err error)
		Login(payload request.UserLogin) (res response.UserLogin, err error)
		Update(id uuid.UUID, payload request.UserUpdate) (res response.UserUpdate, err error)
		Delete(id uuid.UUID) (res response.UserDelete, err error)
	}

	UserRepository interface {
		Register(user *User) (err error)
		Login(user *User) (err error)
		Update(user *User) (err error)
		Delete(id uuid.UUID) (err error)
	}
)
