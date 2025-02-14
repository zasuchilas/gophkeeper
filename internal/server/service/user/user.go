package user

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
)

//go:generate mockery --name Repository
type Repository interface {
	CreateUser(ctx context.Context, item *model.User) (int64, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type Service struct {
	cfg  *config.Config
	repo Repository
}

func NewService(cfg *config.Config, repo Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

// TODO: найти где эта проверка уже есть: var _ def.UserService = (*Service)(nil)
