package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upNewSecretsTable, downNewSecretsTable)
}

func upNewSecretsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE SCHEMA IF NOT EXISTS gophkeeper;
		
		CREATE TABLE IF NOT EXISTS gophkeeper.users
		(
			id         SERIAL PRIMARY KEY,
			login      VARCHAR(254)             NOT NULL UNIQUE,
			pass_hash  VARCHAR(254)             NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
		);
		CREATE INDEX IF NOT EXISTS idx_login ON gophkeeper.users (login);
		
		CREATE TABLE IF NOT EXISTS gophkeeper.secrets
		(
			id         BIGSERIAL PRIMARY KEY,
			name       VARCHAR(254) DEFAULT '',
			data       BYTEA     NOT NULL,
			size       INTEGER   NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			user_id    BIGINT    NOT NULL REFERENCES gophkeeper.users (id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downNewSecretsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`
		DROP INDEX IF EXISTS idx_user_id;
		DROP TABLE IF EXISTS gophkeeper.secrets;
		DROP TABLE IF EXISTS gophkeeper.users;
	`)
	if err != nil {
		return err
	}
	return nil
}
