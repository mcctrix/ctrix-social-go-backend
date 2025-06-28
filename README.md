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
