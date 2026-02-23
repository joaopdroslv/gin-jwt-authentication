package auth

import (
	"database/sql"
)

type AuthRepositoryMysql struct {
	db *sql.DB
}

func NewAuthRepositoryMysql(db *sql.DB) *AuthRepositoryMysql {

	return &AuthRepositoryMysql{db: db}
}
