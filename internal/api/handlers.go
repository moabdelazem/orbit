package api

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/moabdelazem/orbit/internal/utils"
)

// HealthResponse represents the health check response structure
type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Hostname  string    `json:"hostname"`
	Version   string    `json:"version"`
}

// Handler holds dependencies for the API handlers
type Handler struct {
	logger *log.Logger
}

// NewHandler creates a new Handler with the given dependencies
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}

// RegisterRoutes sets up the API routes
func RegisterRoutes(router *mux.Router, logger *log.Logger) {
	handler := NewHandler(logger)

	router.HandleFunc("/health", handler.HealthHandler).Methods("GET")
	router.HandleFunc("/api/v1/status", handler.StatusHandler).Methods("GET")

	// Add middleware for logging
	router.Use(loggingMiddleware(logger))
}

// HealthHandler handles health check requests
func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		h.logger.Printf("Failed to get hostname: %v", err)
		hostname = "unknown"
	}

	response := HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Hostname:  hostname,
		Version:   "1.0.0",
	}

	h.logger.Println("Health check performed")
	utils.WriteJSON(w, response, http.StatusOK)
}

// StatusHandler provides detailed system status
func (h *Handler) StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":    "operational",
		"timestamp": time.Now(),
		"services": map[string]string{
			"database": "connected",
			"cache":    "connected",
		},
	}

	utils.WriteJSON(w, status, http.StatusOK)
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(logger *log.Logger) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			logger.Printf("Request started: %s %s", r.Method, r.URL.Path)

			next.ServeHTTP(w, r)

			logger.Printf("Request completed: %s %s in %v", r.Method, r.URL.Path, time.Since(startTime))
		})
	}
}
