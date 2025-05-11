package maps_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wanderin/internal/info_destination/maps_services"
)

type DestinationHandler struct {
	Service *maps_services.DestinationService
}

func (h *DestinationHandler) AddDestination(c *gin.Context) {
	var req struct {
		Name     string  `json:"name"`
		Category string  `json:"category"`
		Location string  `json:"location"`
		Latitude float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Rating   float64 `json:"rating"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.AddDestination(req.Name, req.Category, req.Location, req.Latitude, req.Longitude, req.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add destination"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Destination added successfully"})
}

func (h *DestinationHandler) GetDestinations(c *gin.Context) {
	destinations, err := h.Service.GetDestinations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch destinations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"destinations": destinations})
}