package main

import (
	"log"

	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/database"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database migration completed successfully")
}
