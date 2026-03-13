package service

import "errors"

var (
	ErrInvalidInput = errors.New("invalid_input")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("not_found")
)
