package model

import "time"

type Secret struct {
	ID         int64     `db:"id"`
	Name       string    `db:"name"`
	Data       []byte    `db:"data"`
	Size       int64     `db:"size"`
	SecretType string    `db:"secret_type"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
	UserID     int64     `db:"user_id"`
}

type SecretFilters struct {
	UserID int64 `db:"user_id"`
	Limit  int64 `db:"limit"`
	Offset int64 `db:"offset"`
}
