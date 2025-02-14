package model

import "time"

type User struct {
	ID        string    `db:"id"`
	Login     string    `db:"login"`
	Password  string    `db:"pass_hash"`
	CreatedAt time.Time `db:"created_at"`
}
