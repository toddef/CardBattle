package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func TestHealthEndpoint(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()
	r.HandleFunc("/healthz", HealthHandler).Methods("GET")

	// Create a test request
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response body
	var response HealthResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)
}
