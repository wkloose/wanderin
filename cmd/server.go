package cmd

import (
	"log"
	"github.com/gin-gonic/gin"
	"wanderin/internal/registerlogin/handlers"
	"wanderin/internal/registerlogin/services"
	"wanderin/internal/registerlogin/repositories"
	"wanderin/internal/middleware"
	"wanderin/config"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		log.Println("Incoming request:", c.Request.Method, c.Request.URL)
		c.Next()
	})

	userRepo := &repositories.UserRepository{DB: config.DB}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{AuthService: authService}

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
	router.POST("/auth/google/callback", authHandler.GoogleLoginCallback)

	protected := router.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", authHandler.GetProfile)

	return router
}