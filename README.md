# 🚀 Go Skeleton

[![Go Version](https://img.shields.io/badge/Go-1.23.0-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/your-username/go-skeleton)](https://goreportcard.com/report/github.com/your-username/go-skeleton)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/your-username/go-skeleton)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://www.docker.com/)

A production-ready Go skeleton with clean architecture, dependency injection, and modern development practices. Built with Gin, PostgreSQL, Redis, and comprehensive tooling for fast application development.

## ✨ Features

- 🏗️ **Clean Architecture** - Hexagonal architecture with clear separation of concerns
- 🔧 **Dependency Injection** - Built-in DI container for easy service management
- 🗄️ **Database Support** - PostgreSQL with SQLx and migration management
- 🚀 **REST API** - Gin framework with structured routing and middleware
- 📝 **Structured Logging** - Zap logger with configurable levels and formats
- 🐳 **Docker Ready** - Multi-stage Dockerfile with PostgreSQL and Redis
- 🔍 **Code Quality** - GolangCI-Lint with comprehensive linting rules
- 🧪 **Testing** - Built-in test utilities and coverage reporting
- ⚡ **CLI Commands** - Migration and server management commands
- 🔐 **Configuration** - Viper-based configuration with environment support
- 📊 **Health Checks** - Database and service health monitoring

## 🏗️ Architecture

```
├── cmd/                   # Application entrypoints
│   ├── app/               # Main application setup
│   └── serve.go           # Server command
├── config/                # Configuration management
├── internal/              # Private application code
│   ├── common/            # Shared utilities
│   └── ping/              # Example module (Clean Architecture)
│       ├── adapter/       # External adapters (REST, Repository)
│       └── core/          # Business logic and ports
├── pkg/                   # Public packages
│   ├── cache/             # Caching utilities
│   ├── database/          # Database management
│   ├── errors/            # Error handling
│   └── logger/            # Logging utilities
└── migrations/            # Database migrations
```

## 🚀 Quick Start

### Prerequisites

- Go 1.23.0 or higher
- Docker and Docker Compose
- PostgreSQL (via Docker)
- Redis (via Docker)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/go-skeleton.git
   cd go-skeleton
   ```

2. **Copy configuration**
   ```bash
   cp application.sample.yml application.yml
   ```

3. **Start services with Docker**
   ```bash
   make docker-up
   ```

4. **Run database migrations**
   ```bash
   make migrate
   ```

5. **Start the application**
   ```bash
   make run-server
   ```

The server will be available at `http://localhost:8081`

## 📖 Usage

### Available Commands

```bash
# Development
make build              # Build the application
make run                # Run the application
make run-server         # Start the server
make test               # Run tests
make test-coverage      # Run tests with coverage

# Database
make migrate            # Run all migrations
make migrate-create     # Create new migration
make migrate-up         # Apply migrations
make migrate-down       # Rollback migrations

# Docker
make docker-up          # Start all services
make docker-down        # Stop all services
make docker-rebuild     # Rebuild and restart

# Code Quality
make lint               # Run linter
make fmt                # Format code
make vet                # Run go vet
make quality-checks     # Run all quality checks
```

### CLI Commands

```bash
# Start server
go run main.go server

# Database migrations
go run main.go migrate up
go run main.go migrate down
go run main.go migrate create --name=add_users_table
```

### Configuration

The application uses `application.yml` for configuration. Key settings:

```yaml
# Server Configuration
SERVER_PORT: 8081
READ_TIMEOUT_MS: 2000
WRITE_TIMEOUT_MS: 2000

# Database Configuration
DB_DRIVER: postgres
DB_NAME: go_skeleton
DB_HOST: postgres-db
DB_USER: postgres
DB_PASSWORD: postgres
DB_PORT: 5432

# Redis Configuration
REDIS_HOST: localhost
REDIS_PORT: 6379

# Logging Configuration
LOG_LEVEL: debug
LOG_ENCODING: json
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test
go test ./internal/ping/...
```

## 🐳 Docker

### Development

```bash
# Start all services
make docker-up

# View logs
docker-compose logs -f app

# Stop services
make docker-down
```

### Production Build

```bash
# Build production image
docker build -t go-skeleton .

# Run production container
docker run -p 8080:8080 go-skeleton
```

## 📁 Project Structure

### Clean Architecture Implementation

The project follows Clean Architecture principles:

- **Adapters**: External interfaces (REST handlers, repositories)
- **Core**: Business logic and domain models
- **Ports**: Interfaces defining contracts between layers

### Example Module: Ping

```go
// Core business logic
type PingService interface {
    Ping() string
}

// REST adapter
type PingHandler struct {
    PingService PingService
}

// Repository adapter
type PingRepository struct {
    db *sqlx.DB
}
```

## 🔧 Development

### Adding New Modules

1. Create module structure in `internal/`
2. Implement core business logic
3. Create adapters for external interfaces
4. Register routes in the main router
5. Add to dependency injection container

### Code Quality

The project uses GolangCI-Lint with comprehensive rules:

- Code formatting and imports
- Security checks (gosec)
- Performance optimizations
- Error handling
- SQL injection prevention

### Environment Variables

```bash
# Development
ENVIRONMENT=development

# Database
DB_HOST=localhost
DB_PORT=5432
DB_NAME=go_skeleton
DB_USER=postgres
DB_PASSWORD=postgres

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go coding standards
- Write tests for new features
- Update documentation
- Run quality checks before committing

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Zap Logger](https://github.com/uber-go/zap)
- [SQLx](https://github.com/jmoiron/sqlx)
- [Viper Configuration](https://github.com/spf13/viper)
- [Go Standards Project Layout](https://github.com/golang-standards/project-layout)

## 📞 Support

- 📧 Email: 94dairy.spacers@icloud.com
- 🐛 Issues: [GitHub Issues](https://github.com/agung-ap/go-skeleton/issues)
- 📖 Documentation: [Wiki](https://github.com/agung-ap/go-skeleton/wiki)

---

⭐ **Star this repository if you find it helpful!**