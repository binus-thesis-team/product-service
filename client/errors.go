package client

import "errors"

var (
	ErrNotFound            = errors.New("error not found")
	ErrInternalServerError = errors.New("error internal server")
)
