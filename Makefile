# ================================
# Project variables
# ================================
APP_NAME=api
CMD_DIR=cmd/api
BIN_DIR=bin
TMP_DIR=tmp

GO=go
AIR=air
GOOSE=goose

# ================================
# Default target
# ================================
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make run          - Run app normally"
	@echo "  make dev          - Run app with Air (hot reload)"
	@echo "  make build        - Build binary"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make test         - Run tests"
	@echo "  make test-cover   - Run tests with coverage"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Run basic lint"
	@echo "  make tidy         - Go mod tidy"
	@echo "  make migrate-up   - Apply DB migrations (Goose)"
	@echo "  make migrate-down - Rollback last migration (Goose)"
	@echo "  make migrate-status - Show migration status"
	@echo "  make migrate-create NAME=<name> - Create a new Goose migration"

# ================================
# Run
# ================================
.PHONY: run
run:
	$(GO) run ./$(CMD_DIR)

.PHONY: dev
dev:
	$(AIR)

# ================================
# Build
# ================================
.PHONY: build
build:
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

# ================================
# Clean
# ================================
.PHONY: clean
clean:
	@rm -rf $(BIN_DIR) $(TMP_DIR)

# ================================
# Tests
# ================================
.PHONY: test
test:
	$(GO) test ./... -v

.PHONY: test-cover
test-cover:
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -func=coverage.out

# ================================
# Formatting & linting
# ================================
.PHONY: fmt
fmt:
	$(GO) fmt ./...

.PHONY: lint
lint:
	$(GO) vet ./...

.PHONY: tidy
tidy:
	$(GO) mod tidy
# ================================
# Database Migrations (Goose)
# ================================
.PHONY: migrate-up
migrate-up:
	@echo "Running Goose migrations (up)..."
	$(GOOSE) -dir ./migrations postgres "$$DATABASE_URL" up

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back last Goose migration..."
	$(GOOSE) -dir ./migrations postgres "$$DATABASE_URL" down

.PHONY: migrate-status
migrate-status:
	@echo "Goose migration status:"
	$(GOOSE) -dir ./migrations postgres "$$DATABASE_URL" status

.PHONY: migrate-create
migrate-create:
ifndef NAME
	$(error NAME is required. Example: make migrate-create NAME=create_users_table)
endif
	@echo "Creating new Goose migration: $(NAME)"
	# explicitly specify -dir ./migrations so it always goes there
	$(GOOSE) create -dir ./migrations $(NAME) sql
