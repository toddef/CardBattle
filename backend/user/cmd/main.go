package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/toddef/cardbattle/backend/user/pkg/user"
)

func main() {
	// Get environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8086"
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to database
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository
	repo := user.NewPostgresRepository(db)

	// Create service
	service := user.NewUserService(repo)

	// Create handlers
	handlers := user.NewHandlers(service)

	// Create router
	r := mux.NewRouter()
	handlers.RegisterRoutes(r)

	// Add health check
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	}).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Starting user service on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
