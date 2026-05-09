# API Service Hub

A lightweight, high-performance API service built with Go and Gin framework.

## Features

- Go 1.21 with Gin web framework
- PostgreSQL database
- Redis for caching (optional)
- Docker Compose for easy local development
- Health check endpoint
- Structured logging with zerolog
- Configuration management with Viper

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose
- Make (optional)

## Local Development

### Option 1: Using Docker Compose (Recommended)

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd api-service-hub
   ```

2. Start all services:
   ```bash
   docker-compose up -d
   ```

3. Check health endpoint:
   ```bash
   curl http://localhost:8080/healthz
   ```

4. View logs:
   ```bash
   docker-compose logs -f api-service
   ```

5. Stop services:
   ```bash
   docker-compose down
   ```

### Option 2: Running Locally

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Start PostgreSQL (using Docker):
   ```bash
   docker run -d \
     --name postgres \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=postgres \
     -e POSTGRES_DB=api_service_hub \
     -p 5432:5432 \
     postgres:15-alpine
   ```

3. Run the application:
   ```bash
   go run cmd/api/main.go
   ```

4. Check health endpoint:
   ```bash
   curl http://localhost:8080/healthz
   ```

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── internal/
│   ├── config/                  # Configuration loading
│   ├── handler/                 # HTTP handlers
│   ├── middleware/              # HTTP middleware
│   ├── model/                   # Data models
│   ├── repository/              # Data access layer
│   └── service/                 # Business logic
├── pkg/
│   ├── config/                  # Shared configuration utilities
│   └── logger/                  # Logging utilities
├── migrations/                  # Database migrations
├── config.yaml                  # Configuration file
├── Dockerfile                   # Docker image definition
├── docker-compose.yml           # Docker Compose configuration
└── README.md                    # This file
```

## API Endpoints

### Health Check

```
GET /healthz
```

Response:
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T12:00:00Z"
}
```

## Configuration

The application can be configured via:

1. **config.yaml** file
2. Environment variables (override config.yaml)

Example environment variables:
```bash
export SERVER_HOST=0.0.0.0
export SERVER_PORT=8080
export DATABASE_HOST=localhost
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASSWORD=postgres
export DATABASE_DBNAME=api_service_hub
export LOG_LEVEL=info
```

## Development

### Run tests:
```bash
go test ./...
```

### Build:
```bash
go build -o api-service ./cmd/api
```

### Run with hot reload (using air):
```bash
# Install air
go install github.com/air-verse/air@latest

# Run with hot reload
air
```

## License

MIT License
