package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toddef/cardbattle/backend/pkg/server"
)

const (
	port = 8082
)

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/healthz", server.HealthHandler).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting Deck service on port %d", port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
