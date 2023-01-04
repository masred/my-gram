package request

import "github.com/google/uuid"

type (
	PhotoCreate struct {
		Title    string    `validate:"required" json:"title"`
		Caption  string    `json:"caption"`
		PhotoUrl string    `validate:"required" json:"photo_url"`
		UserID   uuid.UUID `json:"user_id"`
	}

	PhotoUpdate struct {
		Title    string    `validate:"required" json:"title"`
		Caption  string    `json:"caption"`
		PhotoUrl string    `validate:"required" json:"photo_url"`
		UserID   uuid.UUID `json:"user_id"`
	}
)
