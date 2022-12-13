package exception

import (
	"errors"
)

const (
	unauthorizedErrorMessage       = "unauthorized"
	invalidErrorMessage            = "invalid"
	notFoundErrorMessage           = "not found"
	invalidAuthTokenErrorMessage   = "invalid authentication token"
	invalidCredentialsErrorMessage = "invalid credentials"
)

var (
	ErrUnauthorized       = errors.New(unauthorizedErrorMessage)
	ErrInvalid            = errors.New(invalidErrorMessage)
	ErrEntityNotFound     = errors.New(notFoundErrorMessage)
	ErrInvalidAuthToken   = errors.New(invalidAuthTokenErrorMessage)
	ErrInvalidCredentials = errors.New(invalidCredentialsErrorMessage)
)
