package repository

import (
	"wanderin/internal/info_destination/models"

	"gorm.io/gorm"
)

type DestinationRepository struct {
	DB *gorm.DB
}

// **Simpan destinasi baru ke database**
func (r *DestinationRepository) CreateDestination(dest *models.Destination) error {
	return r.DB.Create(dest).Error
}

// **Ambil semua destinasi populer**
func (r *DestinationRepository) GetDestinations() ([]models.Destination, error) {
	var destinations []models.Destination
	err := r.DB.Find(&destinations).Error
	return destinations, err
}