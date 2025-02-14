package pgsql

import (
	"context"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
	_ "github.com/zasuchilas/gophkeeper/db/server/migrations"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"log/slog"
	"os"
	"path"
)

type Repository struct {
	db *sqlx.DB
}

func MustInit(ctx context.Context, cfg config.PostgreSQL) *Repository {

	db, err := sqlx.Open("pgx", cfg.DSN)
	if err != nil {
		slog.Error("opening connection to postgresql", logger.Err(err))
	}

	// migrations
	err = goose.Up(db.DB, path.Join("db", "server", "migrations"))
	if err != nil {
		slog.Error("goose migrations up", logger.Err(err))
		os.Exit(1)
	}
	slog.Info("goose migrations up successful")

	return &Repository{
		db: db,
	}
}
