package errors

import "errors"

var (
	ErrUserAlreadyExist = errors.New("User already exist")
	ErrUserNotFound     = errors.New("User not found")
	InternalServerError = errors.New("Internal server error")
	Forbidden           = errors.New("Forbidden")
	BadRequest          = errors.New("Bad Request")
)
