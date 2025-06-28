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
