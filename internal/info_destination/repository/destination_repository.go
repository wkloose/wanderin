package repository

import (
	"wanderin/internal/info_destination/models1"

	"gorm.io/gorm"
)

type DestinationRepository struct {
	DB *gorm.DB
}

func (r *DestinationRepository) CreateDestination(dest *models1.Destination) error {
	return r.DB.Create(dest).Error
}

func (r *DestinationRepository) GetDestinations() ([]models1.Destination, error) {
	var destinations []models1.Destination
	err := r.DB.Find(&destinations).Error
	return destinations, err
}
