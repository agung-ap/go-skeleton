.PHONY: build run test clean docker-up docker-down migrate
# Default target
default: build

EXECUTABLE_NAME=golang-skeleton

# Development commands
build:
	@mkdir -p bin
	go build -o bin/$(EXECUTABLE_NAME) main.go

run:
	go run main.go

run-server:
	go run main.go server

test:
	go test -v ./...

test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

test-integration:
	./scripts/integration_test_runner.sh

test-integration-ping:
	ENVIRONMENT=test INTEGRATION_TEST=true go test -v -tags=integration ./internal/... -timeout=5m

test-unit:
	go test -v ./...

test-all: test-unit test-integration

# Database commands
migrate-up:
	go run main.go migrate up

migrate-down:
	go run main.go migrate down

migrate-create:
	@if [ -z "$(name)" ]; then echo "Usage: make migrate-create name=migration_name"; exit 1; fi
	go run main.go migrate create --name=$(name)

# Legacy migrate command (for backward compatibility)
migrate: migrate-up

seed:
	go run scripts/seed/main.go

# Docker commands
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-rebuild:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up -d

# Cleanup
clean:
	rm -rf bin/
	go clean
	docker-compose down -v

# Development setup
setup: docker-up
	sleep 10
	make migrate
	make seed

# Code quality
lint:
	golangci-lint run --config=.golangci.yml

fmt:
	go fmt ./...

vet:
	go vet ./...

quality-checks: fmt vet lint

# Production build
build-prod:
	@mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/$(EXECUTABLE_NAME) main.go
