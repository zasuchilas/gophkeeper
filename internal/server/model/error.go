package model

import "fmt"

var (
	ErrBadParams    = fmt.Errorf("not valid parameters")
	ErrNotFound     = fmt.Errorf("not found")
	ErrConflict     = fmt.Errorf("already exists")
	ErrRelations    = fmt.Errorf("there are relations")
	ErrBadLoginPass = fmt.Errorf("wrong login or password")
)
