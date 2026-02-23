package auth

import (
	"context"
	"database/sql"
	authdomain "gin-jwt-authentication/internal/auth/domain"
)

type AuthRepositoryMysql struct {
	db *sql.DB
}

func NewAuthRepositoryMysql(db *sql.DB) *AuthRepositoryMysql {

	return &AuthRepositoryMysql{db: db}
}

func (r *AuthRepositoryMysql) GetUserByEmail(ctx context.Context, email string) (*authdomain.Credential, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			credentials.id,
			credentials.email,
			credentials.password_hash,
			credentials.created_at,
			credentials.updated_at
		FROM main.credentials
		WHERE credentials.email = ?
	`, email)

	var credential authdomain.Credential

	if err := row.Scan(
		&credential.ID,
		&credential.Email,
		&credential.PasswordHash,
		&credential.CreatedAt,
		&credential.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &credential, nil
}
