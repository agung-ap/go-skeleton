# Integration Testing Guide

This document explains how to run integration tests using the `Dockerfile.test` container.

## Overview

The `Dockerfile.test` is specifically designed for running integration tests in a controlled environment that mirrors production. It includes:

- **Go Application** - Your application running in test mode
- **PostgreSQL Database** - Test database with isolated data
- **Redis Cache** - Test cache instance
- **Network Isolation** - All services communicate via container network

## Prerequisites

- Podman installed
- podman-compose installed
- curl (for HTTP tests)
- netcat (for service health checks)

## Quick Start

### Run All Integration Tests

```bash
# Run the complete integration test suite
make test-integration

# Or run the script directly
./scripts/integration_test_runner.sh
```

### Run Only Go Integration Tests

```bash
# Set environment and run Go integration tests
export ENVIRONMENT=test
export INTEGRATION_TEST=true
go test -v -tags=integration ./...
```

## Test Environment

### Services

| Service | Port | Purpose |
|---------|------|---------|
| Go Application | 1996 | Your application in test mode |
| PostgreSQL | 1998 | Test database |
| Redis | 1997 | Test cache |

### Configuration

The test environment uses `test.application.yml` with the following settings:

- Database: `postgres-db-test:5432`
- Redis: `redis-db-test:6379`
- Application: `:1996`

## Test Structure

Integration tests are organized by domain and placed in their respective directories:

```
internal/
├── ping/
│   ├── module_integration_test.go          # Complete module integration tests
│   ├── adapter/
│   │   ├── ping_repo/
│   │   │   └── repository_integration_test.go  # Repository integration tests
│   │   └── rest/
│   │       └── ping_handler_integration_test.go # Handler integration tests
│   └── core/
│       └── service/
│           └── service_integration_test.go      # Service integration tests
```

This structure allows for:
- **Domain-specific testing**: Tests are co-located with the code they test
- **Granular testing**: Test specific layers (repository, service, handler) independently
- **Easy maintenance**: Tests are easier to find and maintain
- **Clear separation**: Integration tests are clearly separated from unit tests

## Writing Integration Tests

### Go Integration Tests

Create test files with the `integration` build tag in the appropriate domain directory:

```go
//go:build integration
// +build integration

package ping_test

import (
    "testing"
    "os"
)

func TestMyIntegration(t *testing.T) {
    // Skip if not running integration tests
    if os.Getenv("INTEGRATION_TEST") != "true" {
        t.Skip("Skipping integration test")
    }
    
    // Your integration test code here
}
```

### HTTP Tests

The integration test runner includes HTTP endpoint testing:

```bash
# Test ping endpoint
curl http://localhost:1996/ping

# Expected response:
# {"message":"success","data":{"ping_message":"ping from repository"}}
```

## Test Categories

### 1. API Endpoint Tests
- Test all HTTP endpoints
- Verify response formats
- Check error handling

### 2. Database Integration Tests
- Test database connectivity
- Verify CRUD operations
- Test transactions

### 3. Cache Integration Tests
- Test Redis connectivity
- Verify cache operations
- Test cache invalidation

### 4. Full Stack Tests
- Test complete workflows
- Verify service communication
- Performance testing

## Manual Testing

### Start Test Environment

```bash
# Start all services
podman-compose -f docker-compose.test.yml up -d

# Check status
podman ps

# View logs
podman logs go-skeleton-test
```

### Test Endpoints

```bash
# Health check
curl http://localhost:1996/ping

# Documentation
curl http://localhost:1996/docs/

# Add your custom endpoints here
```

### Stop Test Environment

```bash
# Stop all services
podman-compose -f docker-compose.test.yml down
```

## CI/CD Integration

### GitHub Actions Example

```yaml
name: Integration Tests
on: [push, pull_request]

jobs:
  integration-tests:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      
      - name: Install Podman
        run: |
          sudo apt-get update
          sudo apt-get install -y podman podman-compose
      
      - name: Run Integration Tests
        run: |
          chmod +x scripts/integration_test_runner.sh
          ./scripts/integration_test_runner.sh
```

### GitLab CI Example

```yaml
integration_tests:
  stage: test
  image: docker:latest
  services:
    - docker:dind
  before_script:
    - apk add --no-cache podman podman-compose
  script:
    - chmod +x scripts/integration_test_runner.sh
    - ./scripts/integration_test_runner.sh
```

## Troubleshooting

### Common Issues

1. **Port Already in Use**
   ```bash
   # Check what's using the port
   lsof -i :1996
   
   # Stop conflicting services
   podman-compose -f docker-compose.test.yml down
   ```

2. **Database Connection Failed**
   ```bash
   # Check database logs
   podman logs postgres-db-test
   
   # Verify database is running
   podman exec postgres-db-test psql -U postgres -c "SELECT 1;"
   ```

3. **Application Won't Start**
   ```bash
   # Check application logs
   podman logs go-skeleton-test
   
   # Verify configuration
   podman exec go-skeleton-test cat /app/test.application.yml
   ```

### Debug Mode

Run with verbose output:

```bash
# Enable debug logging
export DEBUG=true
./scripts/integration_test_runner.sh
```

## Best Practices

1. **Isolation**: Each test should be independent
2. **Cleanup**: Always clean up test data
3. **Timeouts**: Set appropriate timeouts for external services
4. **Logging**: Use structured logging for debugging
5. **Parallelization**: Run tests in parallel when possible

## Adding New Tests

1. Create test file in the specific domain directory (e.g., `internal/ping/`)
2. Add `//go:build integration` build tag
3. Use `INTEGRATION_TEST=true` environment check
4. Add test to the integration test runner if needed
5. Update this documentation

## Performance Testing

The integration test environment can also be used for performance testing:

```bash
# Run performance tests
go test -v -tags=integration,performance ./...

# Or use dedicated performance test script
./scripts/performance_test_runner.sh
```

## Security Testing

For security testing, you can add security-specific tests:

```bash
# Run security tests
go test -v -tags=integration,security ./...
```

## Monitoring

Monitor test execution with:

```bash
# Watch container logs
podman logs -f go-skeleton-test

# Monitor resource usage
podman stats

# Check network connectivity
podman network ls
``` 