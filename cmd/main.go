package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/orbit/internal/api"
	"github.com/moabdelazem/orbit/internal/config"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Setup logger
	logger := log.New(os.Stdout, "API: ", log.LstdFlags|log.Lshortfile)
	logger.Printf("Starting server on port %s", cfg.Port)

	// Initialize router
	router := mux.NewRouter()

	// Register routes
	api.RegisterRoutes(router, logger)

	// Start server
	serverAddr := ":" + cfg.Port
	logger.Printf("Server listening on %s", serverAddr)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
