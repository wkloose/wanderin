package repositories

import (
	"fmt"
	"wanderin/internal/registerlogin/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) CreateUser(user models.User) error {
	err := r.DB.Create(&user).Error
	if err != nil {
		fmt.Println("GORM Create error:", err)
	}
	return err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
