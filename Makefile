# Makefile for the A/B Testing Dashboard Project

# Use the current directory name as the default project name
PROJECT_NAME ?= $(shell basename "$(CURDIR)")

# Default command to run when you just type "make"
.DEFAULT_GOAL := help

# ==============================================================================
# Docker Compose Commands
# ==============================================================================

.PHONY: up
up: ## Start all services in the background using Docker Compose
	@echo "🚀 Starting all services..."
	docker-compose up -d --build

.PHONY: down
down: ## Stop and remove all services and the network
	@echo "🛑 Stopping all services..."
	docker-compose down

.PHONY: restart
restart: down up ## Restart all services (equivalent to down then up)
	@echo "🔄 Restarting all services..."

.PHONY: logs
logs: ## View the logs of all running services
	@echo "📜 Tailing logs..."
	docker-compose logs -f

.PHONY: logs-backend
logs-backend: ## View the logs of the backend service only
	@echo "📜 Tailing backend logs..."
	docker-compose logs -f backend

.PHONY: logs-frontend
logs-frontend: ## View the logs of the frontend service only
	@echo "📜 Tailing frontend logs..."
	docker-compose logs -f frontend

.PHONY: ps
ps: ## List all running containers for this project
	@echo "📋 Listing running containers..."
	docker-compose ps

# ==============================================================================
# Development & Testing Commands
# ==============================================================================

.PHONY: test
test: ## Run backend unit tests
	@echo "🧪 Running backend tests..."
	@cd backend && go test ./...

.PHONY: build
build: ## Build all Docker images without starting the services
	@echo "🏗️ Building all Docker images..."
	docker-compose build

.PHONY: build-backend
build-backend: ## Build only the backend Docker image
	@echo "🏗️ Building backend image..."
	docker-compose build backend

.PHONY: build-frontend
build-frontend: ## Build only the frontend Docker image
	@echo "🏗️ Building frontend image..."
	docker-compose build frontend

# ==============================================================================
# Housekeeping Commands
# ==============================================================================

.PHONY: clean
clean: down ## Stop services and remove the persistent database volume
	@echo "🧹 Cleaning up... This will delete the database data!"
	docker-compose down -v

.PHONY: prune
prune: ## Remove all stopped containers and dangling images (Docker system-wide)
	@echo "🗑️ Pruning Docker system..."
	docker system prune -f

# ==============================================================================
# Help
# ==============================================================================

.PHONY: help
help: ## Display this help screen
	@echo "Usage: make [command]"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'