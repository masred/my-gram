package entity

import (
	"time"

	"github.com/google/uuid"
)

type SocialMedia struct {
	ID             uuid.UUID `gorm:"primaryKey"`
	Name           string    `gorm:"type:VARCHAR(255) NOT NULL"`
	SocialMediaUrl string    `gorm:"type:TEXT NOT NULL"`
	UserID         uuid.UUID
	User           User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
