package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/pkg/passhash"
	"log/slog"
)

func (s *Service) Login(ctx context.Context, item *model.User) error {

	user, err := s.repo.GetUserByLogin(ctx, item.Login)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return fmt.Errorf("login failed: %w", model.ErrBadLoginPass)
		}
		slog.Error("login error: %s", err.Error())
		return fmt.Errorf("login failed: something went wrong")
	}

	ok := passhash.CheckPasswordHash(item.Password, user.Password)
	if !ok {
		return fmt.Errorf("login failed: %w", model.ErrBadLoginPass)
	}

	// TODO: set jwt token

	return nil
}
