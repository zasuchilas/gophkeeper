package external

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zasuchilas/gophkeeper/internal/server/api/helper"
	"github.com/zasuchilas/gophkeeper/internal/server/jwtmanager"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	desc "github.com/zasuchilas/gophkeeper/pkg/secretsv1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

//go:generate mockery --name SecretsService
type SecretsService interface {
	List(ctx context.Context, filters *model.SecretFilters) ([]model.Secret, error)
	Get(ctx context.Context, id int64) (*model.Secret, error)
	Create(ctx context.Context, item *model.Secret) (*model.Secret, error)
	Update(ctx context.Context, item *model.Secret) (*model.Secret, error)
	Delete(ctx context.Context, id int64) error
}

type SecretsAPI struct {
	desc.UnimplementedSecretsV1Server
	Service SecretsService
}

func NewSecretsAPI(useCases *service.All) *SecretsAPI {
	return &SecretsAPI{
		Service: useCases.Secrets,
	}
}

func (i *SecretsAPI) List(ctx context.Context, in *desc.ListSecretsRequest) (*desc.ListSecretsResponse, error) {

	claims, err := jwtmanager.GetClaims(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	slog.Info("REQUEST (api level logging)", slog.Int64("UserID", claims.ID))
	return nil, nil
}

func (i *SecretsAPI) Get(ctx context.Context, in *desc.SecretRequest) (*desc.GetSecretResponse, error) {

	return &desc.GetSecretResponse{
		UpdatedAt:  timestamppb.New(time.Now()),
		FakeName:   "fake",
		Size:       333,
		SecretType: desc.SecretType_UNKNOWN,
		Data:       nil,
	}, nil
}

func (i *SecretsAPI) Create(ctx context.Context, in *desc.CreateSecretRequest) (*desc.CreateSecretResponse, error) {

	return nil, nil
}

func (i *SecretsAPI) Update(ctx context.Context, in *desc.UpdateSecretRequest) (*desc.UpdateSecretResponse, error) {

	return nil, nil
}

func (i *SecretsAPI) Delete(ctx context.Context, in *desc.SecretRequest) (*empty.Empty, error) {

	return nil, nil
}
