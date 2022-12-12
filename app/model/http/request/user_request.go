package request

type (
	UserRegister struct {
		Username string `validate:"required"  json:"username"`
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required,min=6" json:"password"`
		Age      int    `validate:"required,gt=8" json:"age"`
	}

	UserLogin struct {
		Email    string `validate:"required" json:"email"`
		Password string `validate:"required" json:"password"`
	}

	UserUpdate struct {
		Email    string `validate:"required,email" json:"email"`
		Username string `validate:"required" json:"username"`
	}
)
