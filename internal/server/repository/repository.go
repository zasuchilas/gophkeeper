package repository

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrAlreadyExists = errors.New("already exists")
)

// UserRepository _
type UserRepository interface {
}

// SecretsRepository _
type SecretsRepository interface {
}
