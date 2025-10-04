package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/database"
	"github.com/your-username/coffee-cups-system/internal/logger"
	"github.com/your-username/coffee-cups-system/internal/server"
	"github.com/your-username/coffee-cups-system/internal/services"
	"github.com/your-username/coffee-cups-system/internal/telegram"
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

	// Initialize Telegram bot
	bot, err := telegram.New(cfg.Telegram, services, logger)
	if err != nil {
		logger.Fatal("Failed to initialize Telegram bot", "error", err)
	}

	// Initialize HTTP server
	httpServer := server.New(cfg.Server, services, logger)

	// Start services
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start Telegram bot
	go func() {
		if err := bot.Start(ctx); err != nil {
			logger.Error("Telegram bot error", "error", err)
		}
	}()

	// Start HTTP server
	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Error("HTTP server error", "error", err)
		}
	}()

	logger.Info("Coffee cups system started successfully")

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown
	cancel()
	httpServer.Stop()
}
