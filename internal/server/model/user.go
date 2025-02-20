package model

import "time"

type User struct {
	ID        int64     `db:"id"`
	Login     string    `db:"login"`
	Password  string    `db:"pass_hash"`
	CreatedAt time.Time `db:"created_at"`
}
