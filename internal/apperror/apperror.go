package apperror

import "errors"

var (
	ErrMissingFields       = errors.New("missing fields")
	ErrInvalidCredentials  = errors.New("invalid credentials")
	ErrInvalidFields       = errors.New("invalid fields")
	ErrEmailExists         = errors.New("email already exists")
	ErrUserNotFound        = errors.New("user not found")
	ErrInternalServerError = errors.New("internal server error")
)
