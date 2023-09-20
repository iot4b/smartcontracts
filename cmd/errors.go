package cmd

import "github.com/pkg/errors"

var (
	ErrIsRequired   = errors.New("field is required")
	ErrIsEmpty      = errors.New("value is empty")
	ErrNotSpecified = errors.New("value not specified")
)
