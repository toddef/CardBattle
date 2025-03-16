package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/toddef/cardbattle/backend/auth/pkg/auth"
)

func main() {
	// Get environment variables
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURI := os.Getenv("GOOGLE_REDIRECT_URI")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize Google OAuth client
	googleClient := auth.NewGoogleOAuthClient(clientID, clientSecret, redirectURI)

	// Create auth service
	authService := auth.NewAuthService(googleClient)

	// Set up routes
	http.HandleFunc("/oauth/google", authService.HandleGoogleOAuth)

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
