package main

import (
	authhandler "gin-jwt-authentication/internal/auth/handler"
	authrepository "gin-jwt-authentication/internal/auth/repository"
	authservice "gin-jwt-authentication/internal/auth/service"
	"gin-jwt-authentication/internal/config"
	"gin-jwt-authentication/internal/database"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	env := config.NewEnv()

	db, err := database.NewMysql(env.DockerDatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	authRepository := authrepository.NewAuthRepositoryMysql(db)
	authService := authservice.NewAuthService(authRepository, env.JWTSecret, env.JWTTTL)
	authHandler := authhandler.NewAuthHandler(authService)

	apiV1Group := r.Group("/api/v1")
	authGroup := apiV1Group.Group("/auth")
	authhandler.RegisterRoutes(authGroup, authHandler)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "Ok"})
	})

	r.Run(":" + env.HTTPPort)
}
