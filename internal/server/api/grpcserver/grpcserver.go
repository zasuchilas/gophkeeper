package grpcserver

import (
	"github.com/zasuchilas/gophkeeper/internal/server/api/external"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/jwtmanager"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/zasuchilas/gophkeeper/pkg/secretsv1"
	"github.com/zasuchilas/gophkeeper/pkg/userv1"
)

// Server _
type Server struct {
	cfg        *config.Config
	useCases   *service.All
	grpcServer *grpc.Server
	jwtManager jwtmanager.JWTManager
}

// New _
func New(cfg *config.Config, useCases *service.All, jwtManager jwtmanager.JWTManager) *Server {

	return &Server{
		cfg:        cfg,
		useCases:   useCases,
		jwtManager: jwtManager,
	}
}

// Run _
func (s *Server) Run() {

	s.grpcServer = grpc.NewServer(s.interceptors()...)

	reflection.Register(s.grpcServer)

	userv1.RegisterUserV1Server(s.grpcServer, external.NewUserAPI(s.useCases))
	secretsv1.RegisterSecretsV1Server(s.grpcServer, external.NewSecretsAPI(s.useCases))

	slog.Info("gRPC server is running", slog.String("address", s.cfg.GRPCServer.Address))

	list, err := net.Listen("tcp", s.cfg.GRPCServer.Address)
	if err != nil {
		slog.Error("gRPC listen", logger.Err(err))
		os.Exit(1)
	}

	err = s.grpcServer.Serve(list)
	if err != nil {
		slog.Error("gRPC serve", logger.Err(err))
		os.Exit(1)
	}
}
