package app

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"github.com/zasuchilas/gophkeeper/internal/server/repository"
	"github.com/zasuchilas/gophkeeper/internal/server/repository/pgsql"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type app struct {
	build *buildInfo
	ctx   context.Context
	//grpcServer          *grpcserver.Server
	serverRepo repository.ServerRepository
}

func New(buildVersion, buildDate, buildCommit string) *app {
	return &app{
		build: NewBuildInfo(buildVersion, buildDate, buildCommit),
		ctx:   context.Background(),
	}
}

func (a *app) Run() {

	// loading configuration
	cfg := config.MustLoad()

	// setting logger
	logger.SetupLogger(cfg.Env)
	slog.Info("starting gothkeeper server", slog.String("env", cfg.Env))

	// build info output
	a.buildInfoOutput()

	// repository
	a.serverRepo = pgsql.MustInit(cfg.PostgreSQL)

	slog.Info("OK")

	// shortener service
	//shortenerService := shortener.NewService(a.shortenerRepo, a.secure)

	// grpc server
	//a.grpcServer = grpcserver.NewServer(shortenerService)
	//go a.grpcServer.Run()

	// graceful shutdown
	a.initGracefulShutdown()
}

func (a *app) initGracefulShutdown() {
	idleConnsClosed := make(chan struct{})
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-sigint
		slog.Info("the stop signal has been received", slog.String("signal", sig.String()))
		//a.grpcServer.Stop()
		close(idleConnsClosed)
	}()
	// blocked until the stop signal
	<-idleConnsClosed
	// stopping services
	//a.shortenerRepo.Stop()
	// fin.
	slog.Info("gophkeeper server stopped")
}
