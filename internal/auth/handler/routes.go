package auth

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.RouterGroup, handler *AuthHandler) {

	r.POST("/login", handler.LoginUser)
}
