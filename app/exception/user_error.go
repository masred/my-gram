package exception

import (
	"errors"
)

const (
	usernameAlreadyRegisteredErrorMessage = "username already registered"
	emailAlreadyRegisteredErrorMessage    = "email already registered"
)

var (
	ErrEmailAlreadyRegistered    error = errors.New(emailAlreadyRegisteredErrorMessage)
	ErrUsernameAlreadyRegistered error = errors.New(usernameAlreadyRegisteredErrorMessage)
)
