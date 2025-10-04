package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/database"
	"github.com/your-username/coffee-cups-system/internal/logger"
	"github.com/your-username/coffee-cups-system/internal/server"
	"github.com/your-username/coffee-cups-system/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger := logger.New(cfg.LogLevel)

	// Initialize database
	db, err := database.New(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", "error", err)
	}

	// Initialize services
	services := services.NewServices(db.DB, logger)

	// Initialize HTTP server (without Telegram bot)
	httpServer := server.New(cfg.Server, services, logger)

	// Start HTTP server
	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Error("HTTP server error", "error", err)
		}
	}()

	logger.Info("Coffee cups system started successfully (HTTP only)")
	logger.Info("HTTP server running on http://localhost:8080")
	logger.Info("API endpoints available at http://localhost:8080/api/v1/")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown
	httpServer.Stop()
}
