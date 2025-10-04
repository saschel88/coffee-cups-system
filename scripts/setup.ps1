# Coffee Cups System Setup Script for Windows

Write-Host "🚀 Setting up Coffee Cups System..." -ForegroundColor Green

# Check if Go is installed
try {
    $goVersion = go version
    Write-Host "✅ Go version: $goVersion" -ForegroundColor Green
} catch {
    Write-Host "❌ Go is not installed. Please install Go 1.21 or later." -ForegroundColor Red
    exit 1
}

# Install dependencies
Write-Host "📦 Installing dependencies..." -ForegroundColor Yellow
go mod download
go mod tidy

# Create necessary directories
Write-Host "📁 Creating directories..." -ForegroundColor Yellow
New-Item -ItemType Directory -Force -Path "bin" | Out-Null
New-Item -ItemType Directory -Force -Path "logs" | Out-Null

# Set up environment file if it doesn't exist
if (-not (Test-Path ".env")) {
    Write-Host "📝 Creating .env file..." -ForegroundColor Yellow
    @"
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
"@ | Out-File -FilePath ".env" -Encoding UTF8
    Write-Host "⚠️  Please update the .env file with your actual configuration values." -ForegroundColor Yellow
}

# Build the application
Write-Host "🔨 Building application..." -ForegroundColor Yellow
go build -o bin/coffee-cups-system.exe ./cmd/server

Write-Host "✅ Setup complete!" -ForegroundColor Green
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Cyan
Write-Host "1. Update the .env file with your configuration" -ForegroundColor White
Write-Host "2. Set up your PostgreSQL database" -ForegroundColor White
Write-Host "3. Run 'go run ./cmd/server' to start the application" -ForegroundColor White
Write-Host "4. Or use 'docker-compose up' to run with Docker" -ForegroundColor White
