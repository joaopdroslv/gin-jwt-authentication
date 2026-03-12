package main

import (
	authhandler "gin-jwt-authentication/internal/auth/handler"
	authmiddleware "gin-jwt-authentication/internal/auth/middleware"
	authrepository "gin-jwt-authentication/internal/auth/repository"
	authservice "gin-jwt-authentication/internal/auth/service"
	"gin-jwt-authentication/internal/config"
	"gin-jwt-authentication/internal/infra"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	env := config.NewEnv()

	db, err := infra.NewMysqlDatabase(env.DockerDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	jwtMiddleware := authmiddleware.JWTAuthenticationMiddleware(env.JWTSecret)

	authRepository := authrepository.NewAuthRepositoryMysql(db)
	authService := authservice.NewAuthService(authRepository, env.JWTSecret, env.JWTTTL)
	authHandler := authhandler.NewAuthHandler(authService)

	apiV1Group := r.Group("/api/v1")

	// Public (health check)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Public
	authGroup := apiV1Group.Group("/auth")
	authhandler.RegisterRoutes(authGroup, authHandler)

	// Protected (requires authentication)
	financialGroup := apiV1Group.Group("/financial")
	financialGroup.Use(jwtMiddleware)
	financialGroup.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "you should only see this message after authentication"})
	})

	r.Run(":" + env.HTTPPort)
}
