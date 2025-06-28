#!/bin/bash

# Migration script for improving codebase structure
# This script will help you migrate from the current flat structure to a better organized one

set -e

echo "🚀 Starting codebase structure migration..."

# Create new directory structure
echo "📁 Creating new directory structure..."

# Create main directories
mkdir -p cmd/server
mkdir -p internal/api/handlers/{auth,posts,users,comments,bookmarks,feed}
mkdir -p internal/api/middleware
mkdir -p internal/api/routes
mkdir -p internal/config
mkdir -p internal/domain/{models,repositories,services}
mkdir -p internal/infrastructure/{database/{migrations,repositories},cache,storage,external/{email,notifications}}
mkdir -p internal/pkg/{auth,logger,errors,utils,constants}
mkdir -p pkg
mkdir -p api/docs
mkdir -p scripts
mkdir -p deployments/{docker,kubernetes}
mkdir -p tests/{unit/{handlers,services,repositories},integration/api,fixtures}
mkdir -p docs

echo "✅ Directory structure created"

# Move existing files to new locations
echo "📦 Moving existing files..."

# Move main.go to cmd/server/
if [ -f "main.go" ]; then
    cp main.go cmd/server/main.go
    echo "  ✅ Moved main.go to cmd/server/"
fi

# Move controllers to handlers
if [ -d "controllers" ]; then
    cp controllers/authController.go internal/api/handlers/auth/handler.go
    cp controllers/postManagementController.go internal/api/handlers/posts/handler.go
    cp controllers/userManagementController.go internal/api/handlers/users/handler.go
    cp controllers/commentController.go internal/api/handlers/comments/handler.go
    cp controllers/bookmarkController.go internal/api/handlers/bookmarks/handler.go
    cp controllers/feedController.go internal/api/handlers/feed/handler.go
    echo "  ✅ Moved controllers to handlers/"
fi

# Move routes
if [ -d "routes" ]; then
    cp routes/authRouter.go internal/api/routes/auth.go
    cp routes/postManagementRouter.go internal/api/routes/posts.go
    cp routes/userManagementRouter.go internal/api/routes/users.go
    cp routes/commentRouter.go internal/api/routes/comments.go
    cp routes/bookmarkRouter.go internal/api/routes/bookmarks.go
    cp routes/feedRouter.go internal/api/routes/feed.go
    echo "  ✅ Moved routes/"
fi

# Move middleware
if [ -d "middleware" ]; then
    cp middleware/authMiddleware.go internal/api/middleware/auth.go
    echo "  ✅ Moved middleware/"
fi

# Move models
if [ -d "models" ]; then
    cp models/userModels.go internal/domain/models/user.go
    cp models/postModels.go internal/domain/models/post.go
    echo "  ✅ Moved models/"
fi

# Move database files
if [ -d "db/v1" ]; then
    cp db/v1/postgres.go internal/infrastructure/database/connection.go
    cp db/v1/userManagementDB.go internal/infrastructure/database/repositories/user_repository.go
    cp db/v1/postsDB.go internal/infrastructure/database/repositories/post_repository.go
    cp db/v1/feedDB.go internal/infrastructure/database/repositories/feed_repository.go
    cp db/v1/bookmarkDB.go internal/infrastructure/database/repositories/bookmark_repository.go
    echo "  ✅ Moved database files/"
fi

# Move utils
if [ -d "utils" ]; then
    cp utils/auth.go internal/pkg/auth/jwt.go
    cp utils/util.go internal/pkg/utils/string.go
    echo "  ✅ Moved utils/"
fi

# Move SQL files
if [ -d "sql" ]; then
    cp sql/createTables.sql internal/infrastructure/database/migrations/001_initial_schema.sql
    cp sql/populateDB.sql internal/infrastructure/database/migrations/002_populate_data.sql
    echo "  ✅ Moved SQL files/"
fi

# Move deployment files
if [ -f "docker-compose.yml" ]; then
    cp docker-compose.yml deployments/docker/docker-compose.yml
    echo "  ✅ Moved docker-compose.yml"
fi

echo "✅ File migration completed"

# Create new configuration structure
echo "⚙️  Creating configuration structure..."

cat > internal/config/config.go << 'EOF'
package config

import (
    "os"
    "strconv"
)

type Config struct {
    Server   ServerConfig
    Database DatabaseConfig
    JWT      JWTConfig
    Redis    RedisConfig
}

type ServerConfig struct {
    Port string
    Host string
}

type DatabaseConfig struct {
    Host     string
    Port     int
    User     string
    Password string
    Name     string
    SSLMode  string
}

type JWTConfig struct {
    Secret     string
    Expiration string
}

type RedisConfig struct {
    URL string
}

func Load() (*Config, error) {
    port, _ := strconv.Atoi(getEnv("DB_PORT", "5432"))
    
    return &Config{
        Server: ServerConfig{
            Port: getEnv("PORT", "4000"),
            Host: getEnv("HOST", "localhost"),
        },
        Database: DatabaseConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     port,
            User:     getEnv("DB_USER", "postgres"),
            Password: getEnv("DB_PASSWORD", ""),
            Name:     getEnv("DB_NAME", "ctrix_social"),
            SSLMode:  getEnv("DB_SSLMODE", "disable"),
        },
        JWT: JWTConfig{
            Secret:     getEnv("JWT_SECRET", ""),
            Expiration: getEnv("JWT_EXPIRATION", "24h"),
        },
        Redis: RedisConfig{
            URL: getEnv("REDIS_URL", "redis://localhost:6379"),
        },
    }, nil
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
EOF

echo "✅ Configuration structure created"

# Create Makefile for build commands
echo "🔨 Creating Makefile..."

cat > Makefile << 'EOF'
.PHONY: build run test clean migrate

# Build the application
build:
	go build -o bin/server cmd/server/main.go

# Run the application
run:
	go run cmd/server/main.go

# Run tests
test:
	go test -v ./tests/unit/...
	go test -v ./tests/integration/...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Database migrations
migrate:
	go run cmd/server/main.go migrate

# Reset database
reset-db:
	go run cmd/server/main.go reset

# Initialize database
init-db:
	go run cmd/server/main.go init-db

# Populate database
populate-db:
	go run cmd/server/main.go populate-db

# Docker commands
docker-build:
	docker build -f deployments/docker/Dockerfile -t ctrix-social-backend .

docker-run:
	docker-compose -f deployments/docker/docker-compose.yml up

docker-stop:
	docker-compose -f deployments/docker/docker-compose.yml down

# Development helpers
dev:
	go run cmd/server/main.go

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...
	go vet ./...
EOF

echo "✅ Makefile created"

# Create .env.example
echo "📝 Creating .env.example..."

cat > .env.example << 'EOF'
# Server Configuration
PORT=4000
HOST=localhost

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=ctrix_social
DB_SSLMODE=disable

# JWT Configuration
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRATION=24h

# Redis Configuration
REDIS_URL=redis://localhost:6379

# Cloudinary Configuration (if using)
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

# Email Configuration (if using)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email@gmail.com
SMTP_PASSWORD=your_app_password

# CORS Configuration
ALLOWED_ORIGINS=http://localhost:3000,https://ctrix-social.vercel.app
EOF

echo "✅ .env.example created"

# Create improved README
echo "📖 Creating improved README..."

cat > README.md << 'EOF'
# Ctrix Social Go Backend

A modern, scalable social media backend built with Go and Fiber.

## 🚀 Features

- **Authentication**: JWT-based authentication with refresh tokens
- **User Management**: Profile management, follow/unfollow system
- **Posts**: Create, read, update, delete posts with media support
- **Comments**: Nested comment system with likes
- **Feed**: Personalized feed generation
- **Bookmarks**: Save and manage bookmarked posts
- **Real-time**: WebSocket support for real-time features

## 📁 Project Structure

```
ctrix-social-go-backend/
├── cmd/server/                    # Application entry point
├── internal/                      # Private application code
│   ├── api/                      # HTTP API layer
│   ├── config/                   # Configuration management
│   ├── domain/                   # Business logic and models
│   ├── infrastructure/           # External dependencies
│   └── pkg/                      # Internal shared packages
├── tests/                        # Test files
├── docs/                         # Documentation
└── deployments/                  # Deployment configurations
```

## 🛠️ Tech Stack

- **Framework**: Fiber (Fast HTTP framework)
- **Database**: PostgreSQL
- **Cache**: Redis
- **Authentication**: JWT
- **File Storage**: Cloudinary
- **Testing**: Testify, Testcontainers
- **Documentation**: Swagger/OpenAPI

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- Docker (optional)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/ctrix-social-go-backend.git
   cd ctrix-social-go-backend
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Set up database**
   ```bash
   make init-db
   make populate-db
   ```

5. **Run the application**
   ```bash
   make run
   ```

The server will start on `http://localhost:4000`

### Using Docker

```bash
# Build and run with Docker Compose
make docker-build
make docker-run
```

## 📚 API Documentation

API documentation is available at `/api/docs` when running the server.

### Authentication Endpoints

- `POST /api/auth/signup` - User registration
- `POST /api/auth/login` - User login
- `POST /api/auth/logout` - User logout
- `POST /api/auth/refresh` - Refresh JWT token

### User Management

- `GET /api/profile` - Get user profile
- `PUT /api/profile` - Update user profile
- `GET /api/profile/:id` - Get specific user profile
- `POST /api/profile/follow/:id` - Follow user
- `DELETE /api/profile/follow/:id` - Unfollow user

### Posts

- `GET /api/posts` - Get posts (with pagination)
- `POST /api/posts` - Create new post
- `GET /api/posts/:id` - Get specific post
- `PUT /api/posts/:id` - Update post
- `DELETE /api/posts/:id` - Delete post
- `POST /api/posts/:id/like` - Toggle like on post

### Comments

- `GET /api/comments/:postId` - Get comments for post
- `POST /api/comments` - Create comment
- `PUT /api/comments/:id` - Update comment
- `DELETE /api/comments/:id` - Delete comment
- `POST /api/comments/:id/like` - Toggle like on comment

### Feed

- `GET /api/feed` - Get personalized feed
- `GET /api/feed/trending` - Get trending posts

### Bookmarks

- `GET /api/bookmarks` - Get user bookmarks
- `POST /api/bookmarks/:postId` - Add bookmark
- `DELETE /api/bookmarks/:postId` - Remove bookmark

## 🧪 Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific test categories
go test ./tests/unit/...
go test ./tests/integration/...
```

## 🚀 Deployment

### Production Deployment

1. **Build the application**
   ```bash
   make build
   ```

2. **Set up production environment**
   ```bash
   cp .env.example .env.production
   # Edit .env.production with production values
   ```

3. **Run database migrations**
   ```bash
   make migrate
   ```

4. **Start the server**
   ```bash
   ./bin/server
   ```

### Docker Deployment

```bash
# Build production image
docker build -f deployments/docker/Dockerfile -t ctrix-social-backend:latest .

# Run with Docker Compose
docker-compose -f deployments/docker/docker-compose.yml up -d
```

## 📝 Development

### Code Style

This project follows Go best practices and uses:

- `gofmt` for code formatting
- `golint` for code linting
- `go vet` for code analysis

```bash
# Format code
make fmt

# Lint code
make lint
```

### Adding New Features

1. Create feature branch: `git checkout -b feature/new-feature`
2. Implement the feature following the project structure
3. Add tests for new functionality
4. Update documentation
5. Submit pull request

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues or have questions:

1. Check the [documentation](docs/)
2. Search existing [issues](https://github.com/your-username/ctrix-social-go-backend/issues)
3. Create a new issue with detailed information

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io/) - Fast HTTP framework
- [PostgreSQL](https://www.postgresql.org/) - Database
- [Redis](https://redis.io/) - Caching
- [Cloudinary](https://cloudinary.com/) - Media storage
EOF

echo "✅ README.md created"

# Create Dockerfile
echo "🐳 Creating Dockerfile..."

cat > deployments/docker/Dockerfile << 'EOF'
# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy configuration files
COPY --from=builder /app/.env.example .env

# Expose port
EXPOSE 4000

# Run the application
CMD ["./main"]
EOF

echo "✅ Dockerfile created"

echo "🎉 Migration completed successfully!"
echo ""
echo "📋 Next steps:"
echo "1. Review the new structure in the created directories"
echo "2. Update import paths in moved files"
echo "3. Test that the application still runs: make run"
echo "4. Gradually refactor controllers into handlers and services"
echo "5. Add tests for new structure"
echo "6. Update documentation as needed"
echo ""
echo "🔧 Available commands:"
echo "  make run        - Run the application"
echo "  make test       - Run tests"
echo "  make build      - Build the application"
echo "  make docker-run - Run with Docker"
echo ""
echo "📚 Check the improved README.md for detailed instructions" 