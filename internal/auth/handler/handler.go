package auth

import (
	"errors"
	authschemas "gin-jwt-authentication/internal/auth/schemas"
	authservice "gin-jwt-authentication/internal/auth/service"
	"gin-jwt-authentication/internal/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *authservice.AuthService
}

func NewAuthHandler(authService *authservice.AuthService) *AuthHandler {

	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {

	var body authschemas.RegisterBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	err := h.authService.RegisterUser(c, body)
	if err != nil {
		if errors.Is(err, errs.ErrResourceAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"message": "this email address is already in use"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "sorry, something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registered successfully"})
}

func (h *AuthHandler) LoginUser(c *gin.Context) {

	var body authschemas.LoginBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	token, err := h.authService.LoginUser(c, body)
	if err != nil {
		if errors.Is(err, errs.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		if errs.IsUserStatusRelated(err) {
			c.JSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "sorry, something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}
