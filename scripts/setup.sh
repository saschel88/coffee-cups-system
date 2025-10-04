#!/bin/bash

# Coffee Cups System Setup Script

set -e

echo "🚀 Setting up Coffee Cups System..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | cut -d'o' -f2)
echo "✅ Go version: $GO_VERSION"

# Install dependencies
echo "📦 Installing dependencies..."
go mod download
go mod tidy

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p bin
mkdir -p logs

# Set up environment file if it doesn't exist
if [ ! -f .env ]; then
    echo "📝 Creating .env file..."
    cat > .env << EOF
# Telegram Bot Configuration
TELEGRAM_BOT_TOKEN=your_bot_token_here

# Database Configuration
DATABASE_HOST=localhost
DATABASE_PORT=5432
DATABASE_USER=coffee_user
DATABASE_PASSWORD=coffee_password
DATABASE_DBNAME=coffee_cups
DATABASE_SSLMODE=disable

# Server Configuration
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# Logging
LOG_LEVEL=info
EOF
    echo "⚠️  Please update the .env file with your actual configuration values."
fi

# Build the application
echo "🔨 Building application..."
go build -o bin/coffee-cups-system ./cmd/server

echo "✅ Setup complete!"
echo ""
echo "Next steps:"
echo "1. Update the .env file with your configuration"
echo "2. Set up your PostgreSQL database"
echo "3. Run 'make run' to start the application"
echo "4. Or use 'docker-compose up' to run with Docker"
