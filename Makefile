.PHONY: help build up down logs clean test lint format install-deps

# Default target
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

# Development
install-deps: ## Install all dependencies
	@echo "Installing Go dependencies..."
	@cd services/benchmark-api && go mod tidy
	@cd services/query-service && go mod tidy
	@cd services/metrics-service && go mod tidy
	@echo "Installing Node.js dependencies..."
	@cd web-ui && npm install
	@echo "Dependencies installed successfully!"

# Docker operations
build: ## Build all Docker images
	@echo "Building Docker images..."
	docker-compose build
	@echo "Build completed!"

up: ## Start all services
	@echo "Starting all services..."
	docker-compose up -d
	@echo "Services started! Access the application at:"
	@echo "  Web UI: http://localhost:3000"
	@echo "  API: http://localhost:8080"
	@echo "  Grafana: http://localhost:3001 (admin/admin)"
	@echo "  Prometheus: http://localhost:9090"
	@echo "  MinIO Console: http://localhost:9001 (admin/password)"

down: ## Stop all services
	@echo "Stopping all services..."
	docker-compose down
	@echo "Services stopped!"

down-clean: ## Stop all services and remove volumes
	@echo "Stopping all services and cleaning up..."
	docker-compose down -v --remove-orphans
	@echo "Cleanup completed!"

logs: ## Show logs for all services
	docker-compose logs -f

logs-api: ## Show logs for benchmark API
	docker-compose logs -f benchmark-api

logs-ui: ## Show logs for web UI
	docker-compose logs -f web-ui

# Development mode
dev-infra: ## Start only infrastructure services (for local development)
	@echo "Starting infrastructure services..."
	docker-compose up -d postgres minio hive-metastore trino presto starrocks-fe starrocks-be prometheus grafana
	@echo "Infrastructure services started!"

dev-api: ## Run benchmark API locally
	@echo "Starting benchmark API in development mode..."
	cd services/benchmark-api && go run main.go

dev-ui: ## Run web UI locally
	@echo "Starting web UI in development mode..."
	cd web-ui && npm start

# Testing
test: ## Run all tests
	@echo "Running Go tests..."
	@cd services/benchmark-api && go test ./...
	@cd services/query-service && go test ./...
	@cd services/metrics-service && go test ./...
	@echo "Running Node.js tests..."
	@cd web-ui && npm test -- --coverage --watchAll=false
	@echo "All tests completed!"

test-api: ## Run API tests only
	@cd services/benchmark-api && go test ./...

test-ui: ## Run UI tests only
	@cd web-ui && npm test -- --coverage --watchAll=false

# Code quality
lint: ## Run linting for all services
	@echo "Running Go linting..."
	@cd services/benchmark-api && golangci-lint run
	@cd services/query-service && golangci-lint run
	@cd services/metrics-service && golangci-lint run
	@echo "Running Node.js linting..."
	@cd web-ui && npm run lint
	@echo "Linting completed!"

format: ## Format all code
	@echo "Formatting Go code..."
	@cd services/benchmark-api && go fmt ./...
	@cd services/query-service && go fmt ./...
	@cd services/metrics-service && go fmt ./...
	@echo "Formatting Node.js code..."
	@cd web-ui && npm run format
	@echo "Code formatting completed!"

# Database operations
db-migrate: ## Run database migrations
	@echo "Running database migrations..."
	@cd services/benchmark-api && go run migrations/migrate.go
	@echo "Migrations completed!"

db-seed: ## Seed database with sample data
	@echo "Seeding database with sample data..."
	@cd services/benchmark-api && go run scripts/seed.go
	@echo "Database seeding completed!"

# Monitoring
dashboard: ## Open all application dashboards
	@echo "Opening application dashboards..."
	@open http://localhost:3000 || xdg-open http://localhost:3000 || start http://localhost:3000
	@open http://localhost:3001 || xdg-open http://localhost:3001 || start http://localhost:3001
	@open http://localhost:9090 || xdg-open http://localhost:9090 || start http://localhost:9090

# Cleanup
clean: ## Clean up Docker resources
	@echo "Cleaning up Docker resources..."
	docker system prune -f
	docker volume prune -f
	@echo "Cleanup completed!"

# Documentation
docs: ## Generate API documentation
	@echo "Generating API documentation..."
	@cd services/benchmark-api && swag init
	@echo "Documentation generated! Access it at http://localhost:8080/swagger/"

# Production
deploy: ## Deploy to production (placeholder)
	@echo "Production deployment not configured yet"
	@echo "This would typically deploy to a container orchestration platform"

# Status
status: ## Show status of all services
	@echo "Service Status:"
	@docker-compose ps
