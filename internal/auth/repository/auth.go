package auth

import (
	"context"
	authdomain "gin-jwt-authentication/internal/auth/domain"
)

type AuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*authdomain.Credential, error)
}
