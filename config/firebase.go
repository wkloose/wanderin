package config

import (
	"context"
	"log"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase() (*auth.Client, error) {
	opt := option.WithCredentialsFile("firebase-service-account.json") 
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println("Error initializing Firebase:", err)
		return nil, err
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Println("Error getting Firebase Auth client:", err)
		return nil, err
	}

	return client, nil
}

func ValidateFirebaseToken(idToken string) (*auth.Token, error) {
	client, err := InitFirebase()
	if err != nil {
		return nil, err
	}

	decodedToken, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		log.Println("Invalid Firebase token:", err)
		return nil, err
	}

	return decodedToken, nil
}