package auth

import authservice "gin-jwt-authentication/internal/auth/service"

type AuthHandler struct {
	authService *authservice.AuthService
}
