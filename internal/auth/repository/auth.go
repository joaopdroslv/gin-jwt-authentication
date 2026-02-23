package auth

import (
	"context"
	authmodels "gin-jwt-authentication/internal/auth/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*authmodels.Credential, error)
	RegisterUser(ctx context.Context, user *authmodels.RegistrationData) error
}
