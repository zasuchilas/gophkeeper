package external

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zasuchilas/gophkeeper/internal/server/api/helper"
	"github.com/zasuchilas/gophkeeper/internal/server/converter"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	desc "github.com/zasuchilas/gophkeeper/pkg/secretsv1"
)

//go:generate mockery --name SecretsService
type SecretsService interface {
	List(ctx context.Context, userID int64, filters *model.SecretFilters) ([]*model.Secret, error)
	Create(ctx context.Context, userID int64, item *model.Secret) (*model.Secret, error)
	Get(ctx context.Context, userID, id int64) (*model.Secret, error)
	Update(ctx context.Context, userID int64, item *model.Secret) (*model.Secret, error)
	Delete(ctx context.Context, userID, id int64) error
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

	userID, err := helper.GetCtxUserID(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	items, err := i.Service.List(ctx, userID, &model.SecretFilters{Limit: 100})
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	result := desc.ListSecretsResponse{
		Items: converter.ToSecretListFromService(items),
	}

	return &result, nil
}

func (i *SecretsAPI) Create(ctx context.Context, in *desc.CreateSecretRequest) (*desc.CreateSecretResponse, error) {

	userID, err := helper.GetCtxUserID(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	item := model.Secret{
		Name:       in.Name,
		Data:       in.Data,
		SecretType: in.SecretType.String(),
	}

	next, err := i.Service.Create(ctx, userID, &item)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	updatedAt := helper.TimeToProto(next.UpdatedAt)
	result := desc.CreateSecretResponse{
		Id:        next.ID,
		Name:      next.Name,
		Size:      next.Size,
		UpdatedAt: updatedAt,
	}

	return &result, nil
}

func (i *SecretsAPI) Get(ctx context.Context, in *desc.SecretRequest) (*desc.Secret, error) {

	userID, err := helper.GetCtxUserID(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	item, err := i.Service.Get(ctx, userID, in.Id)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	result := converter.ToSecretFromService(item)

	return result, nil
}

func (i *SecretsAPI) Update(ctx context.Context, in *desc.UpdateSecretRequest) (*desc.Secret, error) {

	userID, err := helper.GetCtxUserID(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	item := model.Secret{
		ID:         in.Id,
		Name:       in.Name,
		Data:       in.Data,
		SecretType: in.SecretType.String(),
	}

	next, err := i.Service.Update(ctx, userID, &item)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	result := converter.ToSecretFromService(next)

	return result, nil
}

func (i *SecretsAPI) Delete(ctx context.Context, in *desc.SecretRequest) (*empty.Empty, error) {

	userID, err := helper.GetCtxUserID(ctx)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	err = i.Service.Delete(ctx, userID, in.Id)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	return nil, nil
}
