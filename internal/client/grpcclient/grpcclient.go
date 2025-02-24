package grpcclient

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/zasuchilas/gophkeeper/internal/client/config"
	"github.com/zasuchilas/gophkeeper/internal/client/model"
	"github.com/zasuchilas/gophkeeper/pkg/secretsv1"
	"github.com/zasuchilas/gophkeeper/pkg/userv1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"time"
)

var ApiService *service

type service struct {
	ctx           context.Context
	conn          *grpc.ClientConn
	userClient    userv1.UserV1Client
	secretsClient secretsv1.SecretsV1Client
	jwt           string
	// TODO: ctx
}

func New(ctx context.Context) {
	conn, err := grpc.NewClient(config.ServerAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(
			fmt.Sprintf("can't connect to gophkeeper server: %s", err.Error()),
		)
	}

	userClient := userv1.NewUserV1Client(conn)
	secretsClient := secretsv1.NewSecretsV1Client(conn)

	ApiService = &service{
		ctx:           ctx,
		conn:          conn,
		userClient:    userClient,
		secretsClient: secretsClient,
	}
}

func (s *service) Stop() {
	err := s.conn.Close()
	if err != nil {
		lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(
			fmt.Sprintf("can't close grpc conn: %s", err.Error()),
		)
	}
}

func (s *service) Register(login, password string) error {

	ctx, cancel := context.WithTimeout(s.ctx, 3*time.Second)
	defer cancel()

	in := userv1.RegisterRequest{
		Login:    login,
		Password: password,
	}
	resp, err := s.userClient.Register(ctx, &in)
	if err != nil {
		return err
	}

	s.writeAccessToken(resp.Jwt)

	return nil
}

func (s *service) Login(login, password string) error {

	ctx, cancel := context.WithTimeout(s.ctx, 3*time.Second)
	defer cancel()

	in := userv1.LoginRequest{
		Login:    login,
		Password: password,
	}
	resp, err := s.userClient.Login(ctx, &in)
	if err != nil {
		return err
	}

	s.writeAccessToken(resp.Jwt)

	return nil
}

func (s *service) GetSecretList() ([]model.ListSecretItem, error) {

	ctx, cancel := context.WithTimeout(s.setAccessTokenInCtx(s.ctx), 3*time.Second)
	defer cancel()

	in := secretsv1.ListSecretsRequest{
		Limit: 100,
	}
	resp, err := s.secretsClient.List(ctx, &in)
	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, nil
	}

	items := make([]model.ListSecretItem, len(resp.Items))
	for i := range resp.Items {
		items[i] = model.ListSecretItem{
			ID:         resp.Items[i].Id,
			Name:       resp.Items[i].Name,
			Size:       resp.Items[i].Size,
			SecretType: resp.Items[i].SecretType.String(),
			UpdatedAt:  resp.Items[i].UpdatedAt.AsTime(),
		}
	}

	return items, nil
}

func (s *service) isAuthorized() bool {
	return s.jwt != ""
}

func (s *service) writeAccessToken(jwt string) {
	s.jwt = jwt
}

func (s *service) setAccessTokenInCtx(ctx context.Context) context.Context {
	md := metadata.New(map[string]string{"authorization": s.jwt})
	return metadata.NewOutgoingContext(ctx, md)
}
