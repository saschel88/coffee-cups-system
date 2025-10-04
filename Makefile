.PHONY: build run test clean docker-build docker-run

# Build the application
build:
	go build -o bin/coffee-cups-system ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out

# Build Docker image
docker-build:
	docker build -t coffee-cups-system .

# Run with Docker Compose
docker-run:
	docker-compose up --build

# Stop Docker Compose
docker-stop:
	docker-compose down

# Run database migrations
migrate:
	go run ./cmd/migrate

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Install dependencies
deps:
	go mod download
	go mod tidy

# Generate mocks (if using mockgen)
mocks:
	mockgen -source=internal/services/user_service.go -destination=internal/mocks/user_service_mock.go
