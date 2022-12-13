package exception

import (
	"errors"
)

const (
	usernameAlreadyRegisteredErrorMessage = "username already registered"
	emailAlreadyRegisteredErrorMessage    = "email already registered"
)

var (
	ErrEmailAlreadyRegistered    = errors.New(emailAlreadyRegisteredErrorMessage)
	ErrUsernameAlreadyRegistered = errors.New(usernameAlreadyRegisteredErrorMessage)
)
