package user

import (
	"context"
	"fmt"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/pkg/passhash"
)

func (s *Service) Register(ctx context.Context, item *model.User) error {
	var err error

	// make password hash
	item.Password, err = passhash.HashPassword(item.Password)
	if err != nil {
		return fmt.Errorf("failed to create a password hash: %w", err)
	}

	userID, err := s.repo.CreateUser(ctx, item)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	_ = userID
	// TODO: set jwt token

	return nil
}
