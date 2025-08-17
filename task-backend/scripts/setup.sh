#!/bin/bash

# Task Management API - Setup Script

set -e

echo "🚀 Setting up Task Management API..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.24+ first."
    exit 1
fi

echo "✅ Go is installed: $(go version)"

# Check if MySQL is available
if ! command -v mysql &> /dev/null; then
    echo "⚠️  MySQL client not found. Please install MySQL client."
fi

# Install dependencies
echo "📦 Installing Go dependencies..."
go mod download
go mod tidy

# Create necessary directories
echo "📁 Creating directories..."
mkdir -p bin/
mkdir -p tmp/
mkdir -p logs/

# Copy environment file
if [ ! -f .env ]; then
    echo "📄 Creating .env file..."
    cp .env.example .env
    echo "⚠️  Please update .env file with your database configuration"
fi

# Copy config file
if [ ! -f configs/config.yaml ]; then
    echo "📄 Creating config.yaml file..."
    cp configs/config.example.yml configs/config.yaml
    echo "⚠️  Please update configs/config.yaml with your configuration"
fi

echo ""
echo "🎉 Setup completed!"
echo ""
echo "Next steps:"
echo "1. Update configs/config.yaml with your database configuration"
echo "2. Start MySQL database: make docker-up (or use your own MySQL)"
echo "3. Run database migrations: make migrate-up"
echo "4. Start development server: make dev"
echo ""
echo "For testing:"
echo "- Install VS Code REST Client extension"
echo "- Use files in http-client/ folder for API testing"
echo ""
echo "Happy coding! 🚀"