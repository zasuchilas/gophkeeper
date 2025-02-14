package secrets

import (
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
)

//go:generate mockery --name Repository
type Repository interface {
	GetSecrets(ctx context.Context, filters *model.SecretFilters) ([]model.Secret, error)
	CreateSecret(ctx context.Context, item *model.Secret) (*model.Secret, error)
	GetSecret(ctx context.Context, id int64) (*model.Secret, error)
	UpdateSecret(ctx context.Context, item *model.Secret) (*model.Secret, error)
	DeleteSecret(ctx context.Context, id int64) error
}

type Service struct {
	cfg  *config.Config
	repo Repository
}

func NewService(cfg *config.Config, repo Repository) *Service {
	return &Service{
		cfg:  cfg,
		repo: repo,
	}
}

func (s *Service) List(ctx context.Context, filters *model.SecretFilters) ([]model.Secret, error) {

	result, err := s.repo.GetSecrets(ctx, filters)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Get(ctx context.Context, id int64) (*model.Secret, error) {
	return s.repo.GetSecret(ctx, id)
}

func (s *Service) Create(ctx context.Context, item *model.Secret) (*model.Secret, error) {

	item.UserID = 1
	item.Size = int64(len(item.Data))
	if item.Name == "" {
		item.Name = gofakeit.CelebrityActor() + ": " + gofakeit.Phrase()
	}

	result, err := s.repo.CreateSecret(ctx, item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Update(ctx context.Context, item *model.Secret) (*model.Secret, error) {

	item.Size = int64(len(item.Data))

	result, err := s.repo.UpdateSecret(ctx, item)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Delete(ctx context.Context, id int64) error {

	err := s.repo.DeleteSecret(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
