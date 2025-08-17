# Task Management API - Makefile

.PHONY: help build run test clean dev migrate-up migrate-down docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  dev          - Run with live reload (requires air)"
	@echo "  test         - Run all tests"
	@echo "  test-unit    - Run unit tests only"
	@echo "  test-int     - Run integration tests only"
	@echo "  clean        - Clean build artifacts"
	@echo "  migrate-up   - Run database migrations up"
	@echo "  migrate-down - Run database migrations down"
	@echo "  docker-up    - Start database with Docker"
	@echo "  docker-down  - Stop database Docker containers"

# Build application
build:
	@echo "Building application..."
	go build -o bin/server ./cmd/server/main.go

# Run application
run: build
	@echo "Starting server..."
	./bin/server

# Development with live reload
dev:
	@echo "Starting development server with live reload..."
	air

# Run all tests
test:
	@echo "Running all tests..."
	go test -v ./tests/...

# Run unit tests only
test-unit:
	@echo "Running unit tests..."
	go test -v ./tests/unit/...

# Run integration tests only
test-int:
	@echo "Running integration tests..."
	go test -v ./tests/integration/...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	rm -rf tmp/

# Database migrations
migrate-up:
	@echo "Running database migrations up..."
	@echo "Please run migrations manually using MySQL client:"
	@echo "mysql -u root -p task_manager < migrations/20250816103214_create_users_table.up.sql"
	@echo "mysql -u root -p task_manager < migrations/20250816103219_create_tasks_table.up.sql"

migrate-down:
	@echo "Running database migrations down..."
	@echo "Please run migrations manually using MySQL client:"
	@echo "mysql -u root -p task_manager < migrations/20250816103219_create_tasks_table.down.sql"
	@echo "mysql -u root -p task_manager < migrations/20250816103214_create_users_table.down.sql"

# Docker commands for database
docker-up:
	@echo "Starting MySQL database with Docker..."
	docker run --name task-mysql \
		-e MYSQL_ROOT_PASSWORD=password \
		-e MYSQL_DATABASE=task_manager \
		-p 3306:3306 \
		-d mysql:8.0

docker-down:
	@echo "Stopping MySQL database..."
	docker stop task-mysql || true
	docker rm task-mysql || true

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run

# Generate documentation
docs:
	@echo "API documentation available at:"
	@echo "http://localhost:8080/docs (when server is running)"

# Setup development environment
setup:
	@echo "Setting up development environment..."
	@echo "1. Install dependencies..."
	make deps
	@echo "2. Start database..."
	make docker-up
	@echo "3. Wait for database to be ready..."
	sleep 10
	@echo "4. Run migrations..."
	make migrate-up
	@echo "5. Setup complete! Run 'make dev' to start development server"