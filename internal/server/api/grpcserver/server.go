package grpcserver

import (
	"github.com/zasuchilas/gophkeeper/internal/server/api/external"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/logger"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/zasuchilas/gophkeeper/pkg/secretsv1"
	"github.com/zasuchilas/gophkeeper/pkg/userv1"
)

// Server _
type Server struct {
	cfg      *config.Config
	useCases *service.All
	server   *grpc.Server
}

// New _
func New(cfg *config.Config, useCases *service.All) *Server {

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(grpcServer)

	userv1.RegisterUserV1Server(grpcServer, external.NewUserAPI(useCases))
	secretsv1.RegisterSecretsV1Server(grpcServer, external.NewSecretsAPI(useCases))

	return &Server{
		cfg:      cfg,
		useCases: useCases,
		server:   grpcServer,
	}
}

// Run _
func (s *Server) Run() {
	slog.Info("gRPC server is running", slog.String("address", s.cfg.GRPCServer.Address))

	list, err := net.Listen("tcp", s.cfg.GRPCServer.Address)
	if err != nil {
		slog.Error("gRPC listen", logger.Err(err))
		os.Exit(1)
	}

	err = s.server.Serve(list)
	if err != nil {
		slog.Error("gRPC serve", logger.Err(err))
		os.Exit(1)
	}
}
