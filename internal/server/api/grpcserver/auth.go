package grpcserver

import (
	"context"
	"github.com/zasuchilas/gophkeeper/internal/server/jwtmanager"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
)

type AuthInterceptor struct {
	jwtManager   jwtmanager.JWTManager
	openRoutes   map[string]bool
	commonRoutes map[string]bool
}

func NewAuthInterceptor(jwtManager jwtmanager.JWTManager) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager:   jwtManager,
		openRoutes:   OpenRoutes(),
		commonRoutes: CommonRoutes(),
	}
}

// Unary creates unary gRPC interceptor for routes authorization.
func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		var err error
		ctx, err = i.authorize(ctx, info.FullMethod, req)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context, method string, req interface{}) (context.Context, error) {
	slog.Debug("authorize", slog.String("method", method))

	// public routes
	open, ok := i.openRoutes[method]
	if ok && open {
		return ctx, nil
	}

	// getting metadata from grpc request
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	// getting authorization header
	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	// verifying auth token and getting the claims
	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return ctx, status.Errorf(codes.Unauthenticated, "auth failed: %s", err.Error())
	}

	// enriching the context with claims
	ctx = context.WithValue(ctx, "claims", claims)

	// common routes for everyone authorized users
	open, ok = i.commonRoutes[method]
	if ok && open {
		return ctx, nil
	}

	return ctx, status.Error(codes.PermissionDenied, "you have no permission to access the method")
}
