.PHONY: build run clean test lint docker-build docker-run help mod-tidy

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
BINARY_NAME=orbit
MAIN_PATH=./cmd/main.go

# Docker parameters
DOCKER_IMAGE=orbit-api
DOCKER_TAG=latest

all: test build

build:
	@echo "Building..."
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_PATH)

run:
	@echo "Running..."
	$(GORUN) $(MAIN_PATH)

run-with-env:
	@echo "Running with .env file..."
	source .env && $(GORUN) $(MAIN_PATH)

clean:
	@echo "Cleaning..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out

test:
	@echo "Running tests..."
	$(GOTEST) -v ./... -cover

test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v ./... -coverprofile=coverage.out
	$(GOCMD) tool cover -html=coverage.out

lint:
	@echo "Linting..."
	$(GOVET) ./...
	$(GOFMT) ./...

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE):$(DOCKER_TAG)

mod-tidy:
	@echo "Tidying modules..."
	$(GOMOD) tidy

mod-download:
	@echo "Downloading modules..."
	$(GOMOD) download

dev:
	@echo "Running in development mode..."
	air -c .air.toml

install-deps:
	@echo "Installing dependencies..."
	$(GOGET) github.com/cosmtrek/air
	$(GOGET) github.com/golangci/golangci-lint/cmd/golangci-lint

apply-pod:
	@echo "Applying pod..."
	kubectl apply -f k8s/pod.yaml

delete-pod:
	@echo "Deleting pod..."
	kubectl delete -f k8s/pod.yaml

help:
	@echo "Available commands:"
	@echo "  make build            - Build the application"
	@echo "  make run              - Run the application"
	@echo "  make run-with-env     - Run the application with .env file"
	@echo "  make clean            - Remove binary and coverage files"
	@echo "  make test             - Run tests"
	@echo "  make test-coverage    - Run tests with coverage report"
	@echo "  make lint             - Run linters"
	@echo "  make docker-build     - Build Docker image"
	@echo "  make docker-run       - Run Docker container"
	@echo "  make mod-tidy         - Tidy Go modules"
	@echo "  make mod-download     - Download Go modules"
	@echo "  make dev              - Run in development mode with hot reload"
	@echo "  make install-deps     - Install development dependencies"
	@echo "  make apply-pod        - Apply pod to Kubernetes cluster"