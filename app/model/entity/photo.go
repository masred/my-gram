package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/masred/my-gram/app/model/http/request"
	"github.com/masred/my-gram/app/model/http/response"
)

type (
	Photo struct {
		ID        uuid.UUID `gorm:"primaryKey"`
		Title     string    `gorm:"type:TEXT NOT NULL"`
		Caption   string    `gorm:"type:TEXT NOT NULL"`
		PhotoUrl  string    `gorm:"type:TEXT NOT NULL"`
		UserID    uuid.UUID
		User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	PhotoService interface {
		Create(payload request.PhotoCreate) (res response.PhotoCreate, err error)
		GetAll() (res []response.PhotoGetAll, err error)
		GetOne(id uuid.UUID) (res response.PhotoGetOne, err error)
		Update(id uuid.UUID, payload request.PhotoUpdate) (res response.PhotoUpdate, err error)
		Delete(id uuid.UUID) (res response.PhotoDelete, err error)
	}

	PhotoRepository interface {
		Create(photo *Photo) (err error)
		GetAll() (photos []Photo, err error)
		GetOne(id uuid.UUID) (photo Photo, err error)
		Update(photo *Photo) (err error)
		Delete(id uuid.UUID) (err error)
	}
)
