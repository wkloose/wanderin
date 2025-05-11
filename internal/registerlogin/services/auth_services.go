package services

import (
	"errors"
	"log"
	"wanderin/internal/registerlogin/models"
	"wanderin/internal/registerlogin/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *repositories.UserRepository
}

func (s *AuthService) Register(user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	err = s.UserRepo.DB.Create(&user).Error
	if err != nil {
		log.Println("Error creating user:", err)
		return errors.New("failed to register user")
	}

	return nil
}

func (s *AuthService) Login(email, password string) (*models.User, error) {
	var user models.User

	err := s.UserRepo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Println("User not found:", err)
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println("Invalid password:", err)
		return nil, errors.New("invalid email or password")
	}

	return &user, nil
}

func (s *AuthService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.UserRepo.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *AuthService) RegisterOAuthUser(email, name string) (*models.User, error) {
	var user models.User
	if err := s.UserRepo.DB.Where("email = ?", email).First(&user).Error; err == nil {
		return &user, nil 
	}

	newUser := models.User{
		Username: name,
		Email:    email,
		Password: "", 
	}

	if err := s.UserRepo.DB.Create(&newUser).Error; err != nil {
		log.Println("Error creating user:", err)
		return nil, err
	}

	return &newUser, nil
}