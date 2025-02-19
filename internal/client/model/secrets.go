package model

import (
	"github.com/zasuchilas/gophkeeper/internal/client/secret"
	"time"
)

type Secret struct {
	ID         int64
	Name       string
	Data       secret.LogoPass // TODO: secret.Secret
	Size       int64
	SecretType string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ListSecretItem struct {
	ID         int64
	Name       string
	Size       int64
	SecretType string
	UpdatedAt  time.Time
}
