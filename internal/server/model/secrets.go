package model

import "time"

type Secret struct {
	ID        int64
	Name      string
	Data      []byte
	Size      int64
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
}

type SecretFilters struct {
	Limit  int64
	Offset int64
}
