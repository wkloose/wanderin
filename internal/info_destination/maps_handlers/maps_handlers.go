package maps_handlers 

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wanderin/internal/info_destination/maps_services"
)

type MapsHandler struct {
	MapsService *maps_services.MapsService
}

func (h *MapsHandler) GetLocation(c *gin.Context) {
	address := c.Query("address")

	location, err := h.MapsService.FetchLocation(address)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"location": location})
}

func (h *MapsHandler) GetNearbyPlaces(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")

	places, err := h.MapsService.FetchNearbyPlaces(lat, lon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"places": places})
}