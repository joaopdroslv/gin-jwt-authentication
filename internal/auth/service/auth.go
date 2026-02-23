package auth

import (
	authrepository "gin-jwt-authentication/internal/auth/repository"
	"time"
)

type AuthService struct {
	authRepository authrepository.AuthRepository
	jwtSecret      []byte
	jwtTTL         time.Duration
}

func New(authRepository authrepository.AuthRepository, jwtSecret string, jwtTTL int64) *AuthService {

	return &AuthService{
		authRepository: authRepository,
		jwtSecret:      []byte(jwtSecret),
		jwtTTL:         time.Duration(jwtTTL) * time.Second,
	}
}
