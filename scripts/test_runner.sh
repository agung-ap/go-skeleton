#!/bin/bash

# Go Skeleton Test Runner
# This script runs all unit tests in the project

set -e

echo "🚀 Running Go Skeleton Unit Tests"
echo "================================="

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
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

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_error "Go is not installed. Please install Go to run tests."
    exit 1
fi

print_status "Go version: $(go version)"

# Run tests with coverage
print_status "Running unit tests with coverage..."

# Create coverage directory if it doesn't exist
mkdir -p coverage

# Run tests for all packages
go test -v -race -coverprofile=coverage/coverage.out ./...

# Check if tests passed
if [ $? -eq 0 ]; then
    print_status "All tests passed! ✅"
    
    # Generate coverage report
    print_status "Generating coverage report..."
    go tool cover -html=coverage/coverage.out -o coverage/coverage.html
    
    # Show coverage summary
    print_status "Coverage summary:"
    go tool cover -func=coverage/coverage.out | tail -1
    
    print_status "Coverage report generated: coverage/coverage.html"
else
    print_error "Some tests failed! ❌"
    exit 1
fi

# Run linting if golangci-lint is available
if command -v golangci-lint &> /dev/null; then
    print_status "Running linter..."
    golangci-lint run --config=.golangci.yml
    
    if [ $? -eq 0 ]; then
        print_status "Linting passed! ✅"
    else
        print_warning "Linting issues found. Please check the output above."
    fi
else
    print_warning "golangci-lint not found. Skipping linting."
fi

echo ""
print_status "Test run completed!"
print_status "To view coverage report, open: coverage/coverage.html" 