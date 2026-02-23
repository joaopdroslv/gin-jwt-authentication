package auth

import (
	"context"
	"database/sql"
	authmodels "gin-jwt-authentication/internal/auth/models"
)

type AuthRepositoryMysql struct {
	db *sql.DB
}

func NewAuthRepositoryMysql(db *sql.DB) *AuthRepositoryMysql {

	return &AuthRepositoryMysql{db: db}
}

func (r *AuthRepositoryMysql) GetUserByEmail(ctx context.Context, email string) (*authmodels.Credential, error) {

	row := r.db.QueryRowContext(ctx, `
		SELECT
			credentials.id,
			credentials.email,
			credentials.password_hash,
			credentials.created_at,
			credentials.updated_at,
			users.id,
			users.user_status_id
		FROM main.credentials
		JOIN main.users ON users.user_credential_id = credentials.id
		WHERE credentials.email = ?
	`, email)

	var credential authmodels.Credential

	if err := row.Scan(
		&credential.ID,
		&credential.Email,
		&credential.PasswordHash,
		&credential.CreatedAt,
		&credential.UpdatedAt,
		&credential.UserInfo.ID,
		&credential.UserInfo.UserStatusID,
	); err != nil {
		return nil, err
	}

	return &credential, nil
}

func (r *AuthRepositoryMysql) RegisterUser(ctx context.Context, registrationData *authmodels.RegistrationData) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	res, err := tx.ExecContext(
		ctx,
		`INSERT INTO main.credentials (email, password_hash) VALUES (?, ?)`,
		registrationData.Email,
		registrationData.PasswordHash,
	)
	if err != nil {
		return err
	}

	credentialID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	res, err = tx.ExecContext(
		ctx,
		`
		INSERT INTO main.users (
			user_credential_id,
			user_status_id,
			name,
			birthdate
		) VALUES (?, ?, ? ,?)
		`,
		credentialID,
		registrationData.UserStatusID,
		registrationData.Name,
		registrationData.Birthdate,
	)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
