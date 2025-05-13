package cmd

import (
	"log"
	"github.com/gin-gonic/gin"
	"wanderin/internal/registerlogin/handlers"
	"wanderin/internal/registerlogin/services"
	"wanderin/internal/registerlogin/repositories"
	"wanderin/internal/middleware"
	"wanderin/config"
	"wanderin/internal/info_destination/maps_handlers"
	"wanderin/internal/info_destination/maps_services"
	"wanderin/internal/info_destination/repository"

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

	infoDestinationHandler := &maps_handlers.MapsHandler{
		MapsService: &maps_services.MapsService{},
	}
	router.GET("/location", infoDestinationHandler.GetLocation)
	router.GET("/places", infoDestinationHandler.GetNearbyPlaces)

	destinationRepo := &repository.DestinationRepository{DB: config.DB}
	destinationService := &maps_services.DestinationService{Repo: destinationRepo}
	destinationHandler := &maps_handlers.DestinationHandler{Service: destinationService}

	router.POST("/destination", destinationHandler.AddDestination)
	router.GET("/destination", destinationHandler.GetDestinations)


	return router
}