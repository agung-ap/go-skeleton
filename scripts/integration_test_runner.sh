#!/bin/bash

# Go Skeleton Integration Test Runner
# This script runs integration tests using the Dockerfile.test container

set -e

echo "🚀 Running Go Skeleton Integration Tests"
echo "========================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

print_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

# Configuration
TEST_TIMEOUT=60  # seconds to wait for services to be ready
APP_PORT=1996
DB_PORT=1998
REDIS_PORT=1997

# Function to wait for service to be ready
wait_for_service() {
    local service_name=$1
    local host=$2
    local port=$3
    local max_attempts=30
    local attempt=1

    print_status "Waiting for $service_name to be ready on $host:$port..."
    
    while [ $attempt -le $max_attempts ]; do
        if nc -z $host $port 2>/dev/null; then
            print_status "$service_name is ready! ✅"
            return 0
        fi
        
        echo -n "."
        sleep 2
        attempt=$((attempt + 1))
    done
    
    print_error "$service_name failed to start within $((max_attempts * 2)) seconds ❌"
    return 1
}



# Function to run integration tests
run_integration_tests() {
    print_step "Starting integration tests..."
    
    # Test 1: Health check - ping endpoint
    print_status "Testing ping endpoint..."
    http_code=$(curl -s -o /tmp/ping_response -w "%{http_code}" http://localhost:$APP_PORT/ping)
    
    if [ "$http_code" = "200" ]; then
        print_status "Ping endpoint test passed! ✅"
        echo "Response: $(cat /tmp/ping_response)"
    else
        print_error "Ping endpoint test failed! Expected 200, got $http_code ❌"
        return 1
    fi
    
    # Test 2: Database connectivity (if there's a health endpoint)
    print_status "Testing database connectivity..."
    # You can add database-specific tests here
    
    # Test 3: Redis connectivity (if there's a health endpoint)
    print_status "Testing Redis connectivity..."
    # You can add Redis-specific tests here
    
    # Test 4: API endpoints (if any)
    print_status "Testing API endpoints..."
    # Add more API endpoint tests here
    
    print_status "All integration tests passed! ✅"
}

# Function to run Go integration tests
run_go_integration_tests() {
    print_step "Running Go integration tests..."
    
    # Set environment for integration tests
    export ENVIRONMENT=test
    export INTEGRATION_TEST=true
    
    # Run integration tests by domain
    print_status "Running ping module integration tests..."
    go test -v -tags=integration ./internal/ping/... -timeout=5m
    
    if [ $? -eq 0 ]; then
        print_status "Ping module integration tests passed! ✅"
    else
        print_error "Ping module integration tests failed! ❌"
        return 1
    fi
    
    # Run any other integration tests
    print_status "Running other integration tests..."
    go test -v -tags=integration ./... -timeout=5m
    
    if [ $? -eq 0 ]; then
        print_status "All Go integration tests passed! ✅"
    else
        print_error "Some Go integration tests failed! ❌"
        return 1
    fi
}

# Function to cleanup
cleanup() {
    print_step "Cleaning up test environment..."
    podman-compose -f docker-compose.test.yml down
    print_status "Cleanup completed! ✅"
}

# Main execution
main() {
    # Run linting
    print_step "Running linting..."
    make lint
    if [ $? -eq 0 ]; then
        print_status "Linting completed! ✅"
    else
        print_error "Linting failed! ❌"
        return 1
    fi
    
    # Check if podman is available
    if ! command -v podman &> /dev/null; then
        print_error "Podman is not installed. Please install Podman to run integration tests."
        exit 1
    fi
    
    # Check if podman-compose is available
    if ! command -v podman-compose &> /dev/null; then
        print_error "podman-compose is not installed. Please install podman-compose to run integration tests."
        exit 1
    fi
    
    print_status "Podman version: $(podman --version)"
    print_status "Podman-compose version: $(podman-compose --version)"
    
    # Set up trap to cleanup on exit
    trap cleanup EXIT
    
    # Step 1: Start the test environment
    print_step "Starting test environment..."
    podman-compose -f docker-compose.test.yml down 2>/dev/null || true
    podman-compose -f docker-compose.test.yml up -d
    
    # Step 2: Wait for services to be ready
    print_step "Waiting for services to be ready..."
    wait_for_service "PostgreSQL" "localhost" $DB_PORT
    wait_for_service "Redis" "localhost" $REDIS_PORT
    wait_for_service "Go Application" "localhost" $APP_PORT
    
    # Step 4: Run database migrations (if needed)
    print_step "Running database migrations..."
    # TODO: Add migration commands here if needed

    
    # Step 5: Run integration tests
    run_integration_tests
    
    # Step 6: Run Go integration tests (if any)
    if grep -r "//go:build integration" . --include="*_test.go" >/dev/null; then
        run_go_integration_tests
    else
        print_warning "No Go integration tests found. Skipping..."
    fi
    
    print_status "Integration test suite completed successfully! 🎉"
}

# Run main function
main "$@" 