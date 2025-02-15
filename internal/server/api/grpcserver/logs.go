package grpcserver

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
	"log/slog"
)

// loggingInterceptor creates slog logging interceptor.
func loggingInterceptor() grpc.UnaryServerInterceptor {
	logger := slog.Default()

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	fn := func(l *slog.Logger) logging.Logger {
		return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l.Log(ctx, slog.Level(lvl), msg, fields...)
		})
	}

	return logging.UnaryServerInterceptor(fn(logger), opts...)
}
