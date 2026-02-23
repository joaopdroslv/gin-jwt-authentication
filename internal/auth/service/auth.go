package auth

import (
	"context"
	authdomain "gin-jwt-authentication/internal/auth/domain"
	authrepository "gin-jwt-authentication/internal/auth/repository"
	authschemas "gin-jwt-authentication/internal/auth/schemas"
	"gin-jwt-authentication/internal/enums"
	"gin-jwt-authentication/internal/errs"
	"log"
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

	log.Println(user)

	err = s.validateUserStatus(user)
	if err != nil {
		return "", err
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

func (s *AuthService) validateUserStatus(user *authdomain.Credential) error {

	switch user.UserInfo.UserStatusID {
	case int64(enums.Inactive):
		return errs.ErrInactiveUser
	case int64(enums.EmailConfirmationPending):
		return errs.ErrUserEmailConfirmationPending
	case int64(enums.PasswordCreationPending):
		return errs.ErrUserPasswordCreationPending
	case int64(enums.Deleted):
		return errs.ErrDeletedUser
	}

	return nil
}
