package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"time"
)

// GetUserByLogin returns the User for the given email.
func (r *Repository) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	var result model.User

	err := sqlx.Get(r.db, &result, `
		SELECT
			*
		FROM
			gophkeeper.users
		WHERE
			login = $1
	`, login)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("select by login=%s failed: %w", login, err)
	}

	return &result, nil
}

// CreateUser creates the given User.
func (r *Repository) CreateUser(ctx context.Context, item *model.User) (int64, error) {
	var userID int64

	now := time.Now().UTC()
	item.CreatedAt = now

	err := sqlx.Get(r.db, &userID, `
		INSERT INTO gophkeeper.users (
		  login,
		  pass_hash,
			created_at
		) values ($1, $2, $3)
		returning id`,
		item.Login,
		item.Password,
		item.CreatedAt,
	)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
