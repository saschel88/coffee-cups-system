package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/handlers"
	"github.com/your-username/coffee-cups-system/internal/services"
)

// Server represents the HTTP server
type Server struct {
	httpServer *http.Server
	services   *services.Services
	logger     interface{}
}

// New creates a new HTTP server
func New(cfg config.ServerConfig, services *services.Services, logger interface{}) *Server {
	router := mux.NewRouter()

	// Initialize handlers
	handlers := handlers.New(services, logger)

	// API routes
	api := router.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	api.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	api.HandleFunc("/boxes", handlers.GetBoxes).Methods("GET")
	api.HandleFunc("/boxes", handlers.CreateBox).Methods("POST")
	api.HandleFunc("/boxes/{id}", handlers.GetBox).Methods("GET")
	api.HandleFunc("/coffee-logs", handlers.GetCoffeeLogs).Methods("GET")
	api.HandleFunc("/coffee-logs", handlers.LogCoffee).Methods("POST")
	api.HandleFunc("/payments", handlers.GetPayments).Methods("GET")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	httpServer := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		services:   services,
		logger:     logger,
	}
}

// Start starts the HTTP server
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// Stop gracefully stops the HTTP server
func (s *Server) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	return s.httpServer.Shutdown(ctx)
}
