package exception

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationMessage(fieldError validator.FieldError) error {
	field := strings.ToLower(fieldError.Field())
	param := strings.ToLower(fieldError.Param())

	switch fieldError.Tag() {
	case "required":
		return fmt.Errorf("%s is required", field)
	case "email":
		return fmt.Errorf("invalid %s address", field)
	case "min":
		return fmt.Errorf("%s must be have at least %s characters long", field, param)
	case "gt":
		return fmt.Errorf("%s should be greater than %s", field, param)
	default:
		return errors.New("unknown error")
	}
}
