package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cardbattle/backend/pkg/server"
	"github.com/gorilla/mux"
)

const (
	port = 8083
)

func main() {
	r := mux.NewRouter()

	// Health check endpoint
	r.HandleFunc("/healthz", server.HealthHandler).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting Card service on port %d", port)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
