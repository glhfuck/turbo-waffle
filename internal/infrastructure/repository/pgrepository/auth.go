package pgrepository

import (
	"github.com/glhfuck/turbo-waffle/internal/domain"
	"github.com/jmoiron/sqlx"
)

type authPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *authPostgres {
	return &authPostgres{db: db}
}

func (au *authPostgres) CreateUser(u domain.User) (int, error) {
	var user_id int

	query := `
	INSERT INTO users (name, username, password_hash)
	VALUES
	($1, $2, $3)
	RETURNING user_id
	`
	row := au.db.QueryRow(query, u.Name, u.Username, u.Password)

	err := row.Scan(&user_id)
	if err != nil {
		return 0, err
	}

	return user_id, nil
}

func (au *authPostgres) GetUser(username, password string) (domain.User, error) {
	var u domain.User

	query := `
	SELECT user_id
	FROM users
	WHERE username=$1 AND password_hash=$2
	`
	err := au.db.Get(&u, query, username, password)

	return u, err
}
