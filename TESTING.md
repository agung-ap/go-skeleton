# Testing Guide

This document provides comprehensive information about the testing strategy and implementation for the Go Skeleton project.

## Overview

The project follows a clean testing strategy where interfaces are only used within test packages (`_test`) for mocking external dependencies. The main codebase uses concrete types to avoid unnecessary abstraction.

## Testing Architecture

### Clean Architecture Testing Strategy

```
├── Unit Tests
│   ├── Core Business Logic (Services)
│   ├── Adapters (Repository, REST Handlers)
│   ├── Infrastructure (Cache, Database, Logger)
│   └── Configuration
├── Test-Only Interfaces & Mocks
│   ├── Defined within `_test` packages
│   ├── Database Interface Mocks (in test files)
│   ├── Cache Interface Mocks (in test files)
│   ├── Service Interface Mocks (in test files)
│   └── Repository Interface Mocks (in test files)
└── Integration Tests (Future)
```

## Test Structure

### Test Interfaces (Defined in Test Packages Only)

Interfaces are only created and used within test files to enable mocking:
- **Database Interface**: Only in repository test files
- **Cache Interface**: Only in repository test files  
- **Service Interface**: Only in handler test files
- **Repository Interface**: Only in service test files

### Unit Tests Coverage

#### Core Business Logic
- ✅ `internal/ping/core/service/service_test.go` - Service layer tests
- ✅ `internal/ping/adapter/ping_repo/repository_test.go` - Repository tests
- ✅ `internal/ping/adapter/ping_repo/model_test.go` - Model conversion tests

#### REST Layer
- ✅ `internal/ping/adapter/rest/ping_handler_test.go` - HTTP handler tests

#### Configuration
- ✅ `config/key_test.go` - Configuration helper function tests
- ✅ `config/db_test.go` - Database configuration tests
- ✅ `config/logger_test.go` - Logger configuration tests

#### Infrastructure
- ✅ `pkg/errors/errors_test.go` - Error handling tests
- ✅ `container/container_test.go` - Dependency injection tests
- ✅ `internal/common/http/response_test.go` - HTTP response utility tests

## Interfaces for Testing

Interfaces are only defined within test packages to enable mocking. The main codebase uses concrete types.

### Repository Interface (Test Only)
```go
// Only exists in service test files
type PingRepository interface {
    Ping(ctx context.Context, resp *domain.Ping) error
}
```

### Service Interface (Test Only) 
```go
// Only exists in handler test files
type PingService interface {
    Ping(ctx context.Context, resp *domain.Ping) error
}
```

### Database Interface (Test Only)
```go
// Only exists in repository test files
type Database interface {
    PingContext(ctx context.Context) error
}
```

### Cache Interface (Test Only)
```go
// Only exists in repository test files  
type Cache interface {
    Ping(ctx context.Context) CacheResult
}
```

## Running Tests

### Quick Test Run
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Comprehensive Test Script
```bash
# Run the complete test suite with coverage and linting
./test_runner.sh
```

### Test Specific Packages
```bash
# Test only the service layer
go test ./internal/ping/core/service/

# Test only the repository layer
go test ./internal/ping/adapter/ping_repo/

# Test only configuration
go test ./config/
```

## Test Coverage

The test suite aims for high coverage across all critical components:

- **Core Business Logic**: 100% coverage for service methods
- **Repository Layer**: Complete coverage including error scenarios
- **Configuration**: Full coverage of configuration loading and validation
- **Error Handling**: Comprehensive error scenario testing

## Mock Usage Examples

### Repository Testing with Database and Cache Mocks
```go
func TestPingRepository_Success(t *testing.T) {
    mockDB := new(mocks.MockDatabase)
    mockCache := new(mocks.MockCache)
    mockCacheResult := new(mocks.MockCacheResult)

    mockDB.On("PingContext", mock.Anything).Return(nil)
    mockCache.On("Ping", mock.Anything).Return(mockCacheResult)
    mockCacheResult.On("Err").Return(nil)

    repo := NewPingRepositoryWithInterfaces(mockDB, mockCache)
    // ... test implementation
}
```

### Service Testing with Repository Mock
```go
func TestPingService_Success(t *testing.T) {
    mockRepo := new(mocks.MockPingRepository)
    mockRepo.On("Ping", mock.Anything, mock.AnythingOfType("*domain.Ping")).Return(nil)
    
    svcCtx := port.NewServiceContext(mockRepo)
    service := NewPingService(svcCtx)
    // ... test implementation
}
```

## Testing Dependencies

The project uses the following testing libraries:

- **testify/assert**: For assertions in tests
- **testify/mock**: For creating mocks
- **testify/suite**: For test suites (if needed)

## Build Constraints

⚠️ **Note**: Some tests may be skipped due to build constraints with the Gin framework. These are related to Go version compatibility issues and don't affect the core testing logic.

## Best Practices

1. **Use Interfaces**: All external dependencies are abstracted behind interfaces
2. **Mock External Dependencies**: Database, cache, and HTTP clients are mocked
3. **Test Error Scenarios**: Both success and failure paths are tested
4. **Isolated Tests**: Each test is independent and doesn't affect others
5. **Clear Test Names**: Test names clearly describe what is being tested
6. **Arrange-Act-Assert**: Tests follow the AAA pattern for clarity

## CI/CD Integration

The test script can be integrated into CI/CD pipelines:

```yaml
# Example GitHub Actions step
- name: Run Tests
  run: |
    chmod +x test_runner.sh
    ./test_runner.sh
```

## Future Enhancements

- [ ] Integration tests with test database
- [ ] End-to-end API tests
- [ ] Performance/benchmark tests
- [ ] Contract testing between layers
- [ ] Test data builders/factories

## Coverage Reports

After running tests, coverage reports are generated:

- **HTML Report**: `coverage/coverage.html`
- **Text Report**: Displayed in terminal during test run

Open the HTML report in a browser to see detailed line-by-line coverage information.

## Troubleshooting

### Common Issues

1. **Gin Import Errors**: Some tests may be skipped due to build constraints
2. **Mock Expectations**: Ensure all mock expectations are properly set and verified
3. **Interface Compliance**: Verify that mocks implement the correct interfaces

### Debug Tips

- Use `go test -v` for verbose output
- Check mock expectations with `mockObject.AssertExpectations(t)`
- Use `t.Log()` for debugging test execution
- Run specific tests with `go test -run TestName` 