package auth

import (
	"context"
	authdto "gin-jwt-authentication/internal/auth/dto"
	authmodels "gin-jwt-authentication/internal/auth/models"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*authmodels.Credential, error)
	RegisterUser(ctx context.Context, user *authdto.RegistrationData) error
}
