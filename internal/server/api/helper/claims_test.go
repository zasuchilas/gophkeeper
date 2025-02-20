package helper

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"testing"
)

func TestGetCtxUserID(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "successful",
			args: args{
				ctx: context.WithValue(context.Background(), "claims", &model.AuthClaims{
					RegisteredClaims: jwt.RegisteredClaims{},
					ID:               3,
				}),
			},
			want:    3,
			wantErr: false,
		},
		{
			name: "nil ctx error",
			args: args{
				ctx: nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "empty ctx error",
			args: args{
				ctx: context.Background(),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "successful",
			args: args{
				ctx: context.WithValue(context.Background(), "claims", &model.AuthClaims{
					RegisteredClaims: jwt.RegisteredClaims{},
					ID:               0,
				}),
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCtxUserID(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCtxUserID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCtxUserID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
