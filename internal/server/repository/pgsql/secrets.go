package pgsql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"time"
)

// GetSecrets returns a slice of Secrets, respecting the given limit and offset.
func (r *Repository) GetSecrets(ctx context.Context, filters *model.SecretFilters) ([]model.Secret, error) {
	var result []model.Secret
	userID := 1

	err := sqlx.Get(r.db, &result, `
		SELECT
			*
		FROM
			gophkeeper.secrets
		WHERE
			user_id = $1
		LIMIT $2
		OFFSET $3
	`, userID, filters.Limit, filters.Offset)

	if err != nil {
		return nil, fmt.Errorf("unable to process select: %w", err)
	}

	return result, nil
}

// GetSecret returns the Secret for the given id.
func (r *Repository) GetSecret(ctx context.Context, id int64) (*model.Secret, error) {
	var result model.Secret

	err := sqlx.Get(r.db, &result, `
		SELECT
			*
		FROM
			gophkeeper.secrets
		WHERE
			id = $1
	`, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("select by id=%d failed: %w", id, err)
	}

	return &result, nil
}

// CreateSecret creates the given Secret.
func (r *Repository) CreateSecret(ctx context.Context, item *model.Secret) (*model.Secret, error) {

	now := time.Now().UTC()
	item.CreatedAt = now
	item.UpdatedAt = now

	err := sqlx.Get(r.db, &item.ID, `
		INSERT INTO gophkeeper.secrets (
		  name,
		  data,
		  size,
			created_at,
			updated_at,
			user_id
		) values ($1, $2, $3, $4, $5, $6)
		returning id`,
		item.Name,
		item.Data,
		item.Size,
		item.CreatedAt,
		item.UpdatedAt,
		item.UserID,
	)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// UpdateSecret updates the given Secret.
func (r *Repository) UpdateSecret(ctx context.Context, item *model.Secret) (*model.Secret, error) {

	item.UpdatedAt = time.Now().UTC()

	res, err := r.db.Exec(`
		UPDATE gophkeeper.secrets
		SET 
			name = $2,
		  data = $3,
		  size = $4,
			updated_at = $5
		WHERE
		  id = $1`,
		item.ID,
		item.Name,
		item.Data,
		item.Size,
		item.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("can't update: %w", err)
	}
	ra, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("can't get rows affected: %w", err)
	}
	if ra == 0 {
		return nil, model.ErrNotFound
	}

	return item, nil
}

// DeleteSecret deletes the Secret record matching the given ID.
func (r *Repository) DeleteSecret(ctx context.Context, id int64) error {
	res, err := r.db.Exec(`
		DELETE FROM
			gophkeeper.secrets
		WHERE
			id = $1
	`, id)
	if err != nil {
		return checkError(err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("can't get rows affected: %w", err)
	}
	if ra == 0 {
		return model.ErrNotFound
	}

	return nil
}
