package usecase

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrDuplicateProduct = errors.New("product already exist")
	ErrPermissionDenied = errors.New("permission denied")
)
