package app

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/server/api/grpcserver"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"github.com/zasuchilas/gophkeeper/internal/server/repository/pgsql"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	"github.com/zasuchilas/gophkeeper/internal/server/service/secrets"
	"github.com/zasuchilas/gophkeeper/internal/server/service/user"
	"log/slog"
	"os/signal"
	"syscall"
)

type app struct {
	build *buildInfo
	ctx   context.Context
	grpc  *grpcserver.Server
}

func New(buildVersion, buildDate, buildCommit string) *app {
	return &app{
		build: NewBuildInfo(buildVersion, buildDate, buildCommit),
		ctx:   context.Background(),
	}
}

func (a *app) Run() {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// loading configuration
	cfg := config.MustLoad()

	// setting logger
	logger.SetupLogger(cfg.Env)
	slog.Info("starting gothkeeper server", slog.String("env", cfg.Env))
	a.buildInfoOutput()

	// repository
	repo := pgsql.MustInit(ctx, cfg.PostgreSQL)
	_ = repo

	//tokenManager := middleware.NewJWTManager(
	//	cfg.Server.ExternalAPI.JWTSecrets,
	//	cfg.Server.ExternalAPI.DefaultSessionTTL,
	//	cache.Roles,
	//)

	// getting services (use cases)
	useCases := service.All{
		User:    user.NewService(cfg, repo),
		Secrets: secrets.NewService(cfg, repo),
	}

	// grpc server
	a.grpc = grpcserver.New(cfg, &useCases)
	go a.grpc.Run()

	slog.Info("the server is running (press CTRL+C to stop)")

	// graceful shutdown
	<-ctx.Done()
	slog.Info("shutting down the server ...")

	// ...

	slog.Info("the server has been gracefully stopped")
}
