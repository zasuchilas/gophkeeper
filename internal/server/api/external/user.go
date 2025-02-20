package external

import (
	"context"
	"errors"
	"github.com/zasuchilas/gophkeeper/internal/server/api/helper"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"github.com/zasuchilas/gophkeeper/internal/server/service"
	desc "github.com/zasuchilas/gophkeeper/pkg/userv1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Register(ctx context.Context, item *model.User) (string, error)
	Login(ctx context.Context, item *model.User) (string, error)
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

func (i *UserAPI) Register(ctx context.Context, in *desc.RegisterRequest) (*desc.RegisterResponse, error) {

	item := model.User{
		Login:    in.Login,
		Password: in.Password,
	}

	jwt, err := i.Service.Register(ctx, &item)
	if err != nil {
		return nil, helper.ErrorToGRPC(err)
	}

	return &desc.RegisterResponse{Jwt: jwt}, nil
}

func (i *UserAPI) Login(ctx context.Context, in *desc.LoginRequest) (*desc.LoginResponse, error) {

	item := model.User{
		Login:    in.Login,
		Password: in.Password,
	}

	jwt, err := i.Service.Login(ctx, &item)
	if err != nil {
		if errors.Is(err, model.ErrBadLoginPass) {
			return nil, status.Errorf(codes.Unauthenticated, "wrong login or password")
		}
		return nil, helper.ErrorToGRPC(err)
	}

	return &desc.LoginResponse{Jwt: jwt}, nil
}
