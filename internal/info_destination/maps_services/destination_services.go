package maps_services

import (
	"wanderin/internal/info_destination/models1"
	"wanderin/internal/info_destination/repository"
)

type DestinationService struct {
	Repo *repository.DestinationRepository
}

// **Tambahkan destinasi populer**
func (s *DestinationService) AddDestination(name, category, location string, lat, lon, rating float64) error {
	destination := &models1.Destination{
		Name:      name,
		Category:  category,
		Location:  location,
		Latitude:  lat,
		Longitude: lon,
		Rating:    rating,
	}

	return s.Repo.CreateDestination(destination)
}

// **Ambil semua destinasi populer**
func (s *DestinationService) GetDestinations() ([]models1.Destination, error) {
	return s.Repo.GetDestinations()
}