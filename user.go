package main

import "time"

type User struct {
	ID            int64     `db:"id" json:"-"`
	Username      string    `db:"username" json:"username"`
	Password_hash string    `db:"password_hash" json:"password"`
	Created_at    time.Time `db:"created_at" json:"created_at"`
}

type Session struct {
	ID          int64     `db:"id" json:"id"`
	Token       string    `db:"token" json:"token"`
	Username_id int64     `db:"username_id" json:"username_id"`
	Created_at  time.Time `db:"created_at" json:"created_at"`
	Expires_at  time.Time `db:"expires_at" json:"expires_at"`
}
