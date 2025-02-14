package external

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	desc "github.com/zasuchilas/gophkeeper/pkg/userv1"
)

type UserService interface {
	Register(ctx context.Context, item *model.User) error
	Login(ctx context.Context, item *model.User) error
}

type UserAPI struct {
	desc.UnimplementedUserV1Server
	Service UserService
}

func NewUserAPI(useCases *service.All) *UserAPI {
	return &UserAPI{
		Service: useCases.User,
	}
}

func (i *UserAPI) Register(ctx context.Context, in *desc.RegisterRequest) (*empty.Empty, error) {

	return nil, nil
}

func (i *UserAPI) Login(ctx context.Context, in *desc.LoginRequest) (*empty.Empty, error) {

	return nil, nil
}
