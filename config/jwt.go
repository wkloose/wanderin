package config

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GetTokenDuration() time.Duration {
	durationStr := os.Getenv("JWT_EXPIRATION")

	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Println("Invalid DURATION_JWT value, using default 24h")
		return 24 * time.Hour
	}
	return duration
}

func GenerateToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(GetTokenDuration()).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	return token, err
}