package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	ID int64 `json:"id"`
}
