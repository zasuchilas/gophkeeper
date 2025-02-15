package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/jwtmanager"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/pkg/passhash"
	"log/slog"
)

//go:generate mockery --name Repository
type Repository interface {
	CreateUser(ctx context.Context, item *model.User) (int64, error)
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
}

type Service struct {
	cfg        *config.Config
	repo       Repository
	jwtManager jwtmanager.JWTManager
}

func NewService(cfg *config.Config, repo Repository, jwtManager jwtmanager.JWTManager) *Service {
	return &Service{
		cfg:        cfg,
		repo:       repo,
		jwtManager: jwtManager,
	}
}

func (s *Service) Register(ctx context.Context, item *model.User) (jwt string, err error) {

	// simple validation
	if len(item.Login) == 0 {
		return "", fmt.Errorf("login is required: %w", model.ErrBadParams)
	}
	if len(item.Password) == 0 {
		return "", fmt.Errorf("password is required: %w", model.ErrBadParams)
	}

	// make password hash
	item.Password, err = passhash.HashPassword(item.Password)
	if err != nil {
		return "", fmt.Errorf("failed to create a password hash: %w", err)
	}

	userID, err := s.repo.CreateUser(ctx, item)
	if err != nil {
		return "", fmt.Errorf("failed to create user: %w", err)
	}

	return s.jwtManager.GenerateUserAccessToken(&model.User{ID: userID})
}

func (s *Service) Login(ctx context.Context, item *model.User) (string, error) {

	user, err := s.repo.GetUserByLogin(ctx, item.Login)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return "", fmt.Errorf("login failed: %w", model.ErrBadLoginPass)
		}
		slog.Error("login error: %s", err.Error())
		return "", fmt.Errorf("login failed: something went wrong")
	}

	ok := passhash.CheckPasswordHash(item.Password, user.Password)
	if !ok {
		return "", fmt.Errorf("login failed: %w", model.ErrBadLoginPass)
	}

	return s.jwtManager.GenerateUserAccessToken(user)
}
