package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cardbattle/backend/pkg/server"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Create a new router
	r := mux.NewRouter()
	r.HandleFunc("/healthz", server.HealthHandler).Methods("GET")

	// Create a test request
	req := httptest.NewRequest("GET", "/healthz", nil)
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Check status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check response body
	var response server.HealthResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "ok", response.Status)
}
