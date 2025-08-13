package errors

import (
	"errors"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmptyPassword      = errors.New("empty password")
	ErrEmptyEmail         = errors.New("empty email")
	ErrTokenExpired       = errors.New("token expired")
	ErrInvalidToken       = errors.New("invalid token")
)
