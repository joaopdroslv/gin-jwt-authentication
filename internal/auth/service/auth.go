package auth

import (
	"context"
	authrepository "gin-jwt-authentication/internal/auth/repository"
	authschemas "gin-jwt-authentication/internal/auth/schemas"
	"gin-jwt-authentication/internal/errs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepository authrepository.AuthRepository
	jwtSecret      []byte
	jwtTTL         time.Duration
}

func NewAuthService(authRepository authrepository.AuthRepository, jwtSecret string, jwtTTL int64) *AuthService {

	return &AuthService{
		authRepository: authRepository,
		jwtSecret:      []byte(jwtSecret),
		jwtTTL:         time.Duration(jwtTTL) * time.Second,
	}
}

func (s *AuthService) LoginUser(ctx context.Context, body authschemas.LoginBody) (string, error) {

	user, err := s.authRepository.GetUserByEmail(ctx, body.Email)
	if err != nil {
		return "", errs.ErrInvalidCredentials
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return "", errs.ErrInvalidCredentials
	}

	claims := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(user.UserInfo.ID, 10),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.jwtTTL)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.jwtSecret)
}
