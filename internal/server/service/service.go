package service

import (
	"github.com/zasuchilas/gophkeeper/internal/server/service/secrets"
	"github.com/zasuchilas/gophkeeper/internal/server/service/user"
)

type All struct {
	User    *user.Service
	Secrets *secrets.Service
}
