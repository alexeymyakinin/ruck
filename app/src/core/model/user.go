package model

type User struct {
	ID       uint64 `db:"id"`
	Email    string `db:"email"`
	Username string `db:"username"`
	Password string
}
