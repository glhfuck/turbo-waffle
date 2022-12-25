package domain

type User struct {
	Id       int    `json:"-"        db:"user_id"`
	Name     string `json:"name"     db:"name"     binding:"required"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" db:"password" binding:"required"`
}
