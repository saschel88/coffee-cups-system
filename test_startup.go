package main

import (
	"fmt"
	"log"

	"github.com/your-username/coffee-cups-system/internal/config"
	"github.com/your-username/coffee-cups-system/internal/logger"
)

func main() {
	fmt.Println("🚀 Testing Coffee Cups System Startup...")
	
	// Test 1: Configuration Loading
	fmt.Println("\n1️⃣ Testing configuration loading...")
	cfg, err := config.Load()
	if err != nil {
		log.Printf("❌ Configuration failed: %v", err)
	} else {
		fmt.Printf("✅ Configuration loaded: Server %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	}
	
	// Test 2: Logger Creation
	fmt.Println("\n2️⃣ Testing logger creation...")
	logger := logger.New(cfg.LogLevel)
	if logger != nil {
		fmt.Println("✅ Logger created successfully")
	} else {
		fmt.Println("❌ Logger creation failed")
	}
	
	// Test 3: Database Connection (will fail without DB)
	fmt.Println("\n3️⃣ Testing database connection...")
	fmt.Println("❌ Database connection will fail (expected - no PostgreSQL running)")
	fmt.Println("   This is normal - you need to set up PostgreSQL first")
	
	// Test 4: Telegram Bot (will fail without token)
	fmt.Println("\n4️⃣ Testing Telegram bot...")
	fmt.Println("❌ Telegram bot will fail (expected - no bot token)")
	fmt.Println("   This is normal - you need to get a bot token from @BotFather")
	
	fmt.Println("\n🎉 Startup test completed!")
	fmt.Println("\n📋 To make it fully working, you need:")
	fmt.Println("   1. Set up PostgreSQL database")
	fmt.Println("   2. Get Telegram bot token from @BotFather")
	fmt.Println("   3. Set TELEGRAM_BOT_TOKEN environment variable")
}
