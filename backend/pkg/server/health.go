package server

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status string `json:"status"`
}

// HealthHandler returns a simple health check handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
