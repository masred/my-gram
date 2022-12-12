package response

import (
	"time"

	"github.com/google/uuid"
)

type (
	UserRegister struct {
		ID       uuid.UUID `json:"id"`
		Email    string    `json:"email"`
		Username string    `json:"username"`
		Age      int       `json:"age"`
	}

	UserLogin struct {
		Token string `json:"token"`
	}

	UserUpdate struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		Username  string    `json:"username"`
		Age       int       `json:"age"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserDelete struct {
		Message string `json:"message"`
	}
)
