# ====================================================================================
# DOCKER COMPOSE
# ====================================================================================

build:
	@echo "Building Docker images..."
	@docker-compose build

up:
	@echo "Starting all services..."
	@docker-compose up -d --build

down:
	@echo "Stopping all services and removing volumes..."
	@docker-compose down --volumes

logs:
	@docker-compose logs -f

logs-app:
	@docker-compose logs -f app

restart-app:
	@echo "Rebuilding and restarting Go application..."
	@docker-compose up -d --build app

# ====================================================================================
# DBMATE
# ====================================================================================

db-up:
	@echo "Running pending migrations..."
	@docker-compose run --rm dbmate up

db-down:
	@echo "Rolling back last migration..."
	@docker-compose run --rm dbmate down

## db-new: ex: make db-new name=create_products_table
db-new:
ifeq ($(name),)
	@echo "Error: 'name' is a required argument."
	@echo "Usage: make db-new name=<migration_name>"
	@exit 1
endif
	@echo "Creating new migration: $(name)"
	@docker-compose run --rm dbmate new $(name)

## db-status: Melihat status dari semua migrasi
db-status:
	@echo "Checking migration status..."
	@docker-compose run --rm dbmate status