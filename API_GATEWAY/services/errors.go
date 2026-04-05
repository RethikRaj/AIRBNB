package services

import "errors"

var (
	ErrInvalidCredentials = errors.New("Invalid Credentials")
	ErrNotFound           = errors.New("Resource not found")
)
