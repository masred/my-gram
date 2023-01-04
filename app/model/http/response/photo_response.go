package response

import (
	"time"

	"github.com/google/uuid"
)

type (
	PhotoCreate struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoUrl  string    `json:"photo_url"`
		UserID    uuid.UUID `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
	}

	PhotoGetOne struct {
		UserID uuid.UUID `json:"user_id"`
	}

	PhotoUserGetAll struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}

	PhotoGetAll struct {
		ID        uuid.UUID       `json:"id"`
		Title     string          `json:"title"`
		Caption   string          `json:"caption"`
		PhotoUrl  string          `json:"photo_url"`
		UserID    uuid.UUID       `json:"user_id"`
		CreatedAt time.Time       `json:"created_at"`
		UpdatedAt time.Time       `json:"updated_at"`
		User      PhotoUserGetAll `json:"user"`
	}

	PhotoUpdate struct {
		ID        uuid.UUID `json:"id"`
		Title     string    `json:"title"`
		Caption   string    `json:"caption"`
		PhotoUrl  string    `json:"photo_url"`
		UserID    uuid.UUID `json:"user_id"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	PhotoDelete struct {
		Message string `json:"message"`
	}
)
