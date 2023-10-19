package cmd

import "github.com/pkg/errors"

var (
	ErrInvalidValue = errors.New("invalid value")
	ErrIsRequired   = errors.New("field is required")
	ErrIsEmpty      = errors.New("value is empty")
	ErrNotSpecified = errors.New("value not specified")
)
