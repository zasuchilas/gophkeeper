package model

import "fmt"

var (
	ErrServerError  = fmt.Errorf("internal server error")
	ErrNoClaims     = fmt.Errorf("no claims in context")
	ErrBadParams    = fmt.Errorf("not valid parameters")
	ErrNotFound     = fmt.Errorf("not found")
	ErrConflict     = fmt.Errorf("already exists")
	ErrRelations    = fmt.Errorf("there are relations")
	ErrBadLoginPass = fmt.Errorf("wrong login or password")
	ErrAccessDenied = fmt.Errorf("access denied")
)
