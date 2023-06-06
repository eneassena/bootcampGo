package domain

type User struct {
	UserID   int    `db:"userid"`
	UserName string `db:"username"`
	Email    string `db:"email"`
}
