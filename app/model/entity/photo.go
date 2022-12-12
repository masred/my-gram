package entity

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Title     string    `gorm:"type:TEXT NOT NULL"`
	Caption   string    `gorm:"type:TEXT NOT NULL"`
	PhotoUrl  string    `gorm:"type:TEXT NOT NULL"`
	UserID    uuid.UUID
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
