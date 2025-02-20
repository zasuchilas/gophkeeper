package app

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/client/config"
	"github.com/zasuchilas/gophkeeper/internal/client/grpcclient"
	"github.com/zasuchilas/gophkeeper/internal/client/tui"
	"os/signal"
	"syscall"
)

type app struct {
	build *buildInfo
}

func New(buildVersion, buildDate, buildCommit string) *app {
	return &app{
		build: NewBuildInfo(buildVersion, buildDate, buildCommit),
	}
}

func (a *app) Run() {

	// config
	config.ParseFlags()
	config.BuildInfo = a.build.String()

	// app context
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	defer func() {
		if grpcclient.ApiService != nil {
			grpcclient.ApiService.Stop()
		}
	}()

	// grpc client service
	grpcclient.New(ctx)

	// tui application
	tui.Run()
}
