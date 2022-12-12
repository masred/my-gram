package entity

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Message   string    `gorm:"type:TEXT NOT NULL"`
	UserID    uuid.UUID
	PhotoID   uuid.UUID
	User      User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Photo     Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
