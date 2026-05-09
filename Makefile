.PHONY: all build test clean lint fmt vet docker-build docker-run docker-stop swagger help

# Variables
BINARY_NAME=api-service
BUILD_DIR=bin
GO_FILES=$(shell find . -name '*.go' -not -path './vendor/*')
DOCKER_IMAGE=api-service-hub
DOCKER_CONTAINER=api-service-hub

all: fmt vet test build

## build: Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

## test: Run tests with coverage
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

## test-coverage: Display test coverage
test-coverage: test
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html

## lint: Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run --timeout=5m

## fmt: Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@goimports -w .

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy

## docker-build: Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE):latest .

## docker-run: Run Docker container
docker-run:
	@echo "Running Docker container..."
	@docker-compose up -d

## docker-stop: Stop Docker containers
docker-stop:
	@echo "Stopping Docker containers..."
	@docker-compose down

## docker-logs: Show Docker logs
docker-logs:
	@docker-compose logs -f

## swagger: Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g cmd/api/main.go -o docs

## run: Run the application locally
run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

## dev: Run in development mode with hot reload (requires air)
dev:
	@air

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Available targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
