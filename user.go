package main

type User struct {
	ID            int64  `db:"id" json:"-"`
	Username      string `db:"username" json:"username"`
	Password_hash string `db:"password_hash" json:"password"`
	Logged_in     bool   `db:"logged_in" json:"-"`
}
