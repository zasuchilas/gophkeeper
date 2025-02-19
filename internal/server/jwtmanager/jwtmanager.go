package jwtmanager

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/zasuchilas/gophkeeper/internal/server/config"
	"github.com/zasuchilas/gophkeeper/internal/server/model"
	"time"
)

type JWTManager interface {
	GenerateUserAccessToken(user *model.User) (jwtToken string, err error)
	Verify(accessToken string) (*model.AuthClaims, error)
}

type jwtManager struct {
	cfg *config.Config
}

func New(cfg *config.Config) *jwtManager {
	return &jwtManager{cfg: cfg}
}

func (j *jwtManager) GenerateUserAccessToken(user *model.User) (jwtToken string, err error) {

	now := time.Now().Unix()
	exp := time.Now().Add(j.cfg.JWT.SessionTTL).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": j.cfg.App,
		"aud": j.cfg.App,
		"nbf": now,
		"exp": exp,
		"sub": "user",
		"id":  user.ID,
	})

	jwtToken, err = token.SignedString([]byte(j.cfg.JWT.Secrets[0]))
	if err != nil {
		return "", fmt.Errorf("can't sign jwt: %w", err)
	}

	return jwtToken, err
}

// Verify checks if the token is valid.
func (j *jwtManager) Verify(accessToken string) (*model.AuthClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	// checking given token with all secrets
	for _, secret := range j.cfg.Secrets {
		token, err = jwt.ParseWithClaims(
			accessToken,
			&model.AuthClaims{},
			func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unexpected token signing method")
				}
				return []byte(secret), nil
			},
		)

		// relevant secret found and token is valid
		if err == nil && token.Valid {
			break
		}
	}

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("bad token: %w", err)
	}

	claims, ok := token.Claims.(*model.AuthClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	sub, err := claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("can't get subject name from token claims %w", err)
	}
	_ = sub

	return claims, nil
}

func GetClaims(ctx context.Context) (*model.AuthClaims, error) {
	if ctx == nil {
		return nil, fmt.Errorf("context is missing")
	}
	claims, ok := ctx.Value("claims").(*model.AuthClaims)
	if !ok {
		return nil, model.ErrNoClaims
	}
	return claims, nil
}
