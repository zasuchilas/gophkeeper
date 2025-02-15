package grpcserver

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *Server) interceptors() []grpc.ServerOption {

	// creating authInterceptor
	authInterceptor := NewAuthInterceptor(s.jwtManager)

	return []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(
			loggingInterceptor(),
			authInterceptor.Unary(),
		),
	}
}
