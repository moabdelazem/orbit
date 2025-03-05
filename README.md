# Orbit API

A simple, well-structured Go API service.

## Project Structure

```
orbit/
├── cmd/
│   └── main.go         # Application entry point
├── internal/
│   ├── api/            # API handlers and routes
│   ├── config/         # Configuration management
│   └── utils/          # Utility functions
├── .env                # Environment variables
├── .air.toml           # Air configuration for hot reloading
├── Dockerfile          # Container definition
├── Makefile            # Build and development commands
└── README.md           # Project documentation
```

## Getting Started

### Running Locally

```bash
# Run the application with default settings
make run

# Run the application with .env file
make run-with-env

# Run with hot reloading during development
make dev
```

### Environment Variables

Copy `.env.example` to `.env` and adjust the values to your needs:

```bash
cp .env.example .env
```

Key environment variables:

- `PORT`: Server port (default: 8080)
- `LOG_LEVEL`: Logging level (default: info)
- `ENV`: Environment name (default: development)
- See the `.env` file for all available configuration options

### Available Make Commands

```bash
# Building and running
make build            # Build the application
make run              # Run the application
make run-with-env     # Run with .env file
make dev              # Run with hot reloading

# Testing and code quality
make test             # Run tests
make test-coverage    # Run tests with coverage report
make lint             # Run linters

# Docker commands
make docker-build     # Build Docker image
make docker-run       # Run Docker container

# Dependency management
make mod-tidy         # Tidy Go modules
make mod-download     # Download Go modules
```

## Building and Running with Docker

```bash
# Build the Docker image
make docker-build

# Run the container
make docker-run
```

## API Endpoints

- `GET /health` - Health check endpoint
- `GET /api/v1/status` - Detailed service status
