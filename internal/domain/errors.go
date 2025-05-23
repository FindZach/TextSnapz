package domain

import "errors"

// Custom errors for the domain layer.
var (
	ErrSnapNotFound = errors.New("snap not found")
	ErrSnapExpired  = errors.New("snap has expired")
)
