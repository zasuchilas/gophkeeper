package app

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/client/config"
	"github.com/zasuchilas/gophkeeper/internal/client/grpcclient"
	"github.com/zasuchilas/gophkeeper/internal/client/tui"
	"os/signal"
	"syscall"
)

type app struct{}

func New() *app {
	return &app{}
}

func (a *app) Run() {

	// config
	config.ParseFlags()

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
