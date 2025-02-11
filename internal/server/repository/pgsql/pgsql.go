package pgsql

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	_ "github.com/zasuchilas/gophkeeper/db/migrations"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"log/slog"
	"os"
	"path"
)

type repository struct {
	db *sql.DB
}

func MustInit(cfg config.PostgreSQL) *repository {

	db, err := sql.Open("pgx", cfg.DSN)
	if err != nil {
		slog.Error("opening connection to postgresql", logger.Err(err))
	}

	// migrations
	err = goose.Up(db, path.Join("db", "migrations"))
	if err != nil {
		slog.Error("goose migrations up", logger.Err(err))
		os.Exit(1)
	}
	slog.Info("goose migrations up successful")

	return &repository{
		db: db,
	}
}
