# Declares targets that are not actual files but just commands.

# Prevents Make from getting confused if a file named build, run, etc. exists in your directory.
.PHONY: build run test clean swagger migrate docker-build docker-run migrate-up migrate-down

# Build the application
# compiles the Go app into bin/server
build:
	go build -o bin/server cmd/server

# Run the application
# runs your app directly without compiling first
run:
	go run ./cmd/server

# Run tests
test:
	go test -v -cover ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Generate swagger documentation
swagger:
	swag init -g cmd/server/main.go -o internal/docs

# Run database migrations
migrate:
	./scripts/migrate.sh

# Clean build files
clean:
	rm -f bin/server
	rm -f coverage.out

# Build Docker image
docker-build:
	docker build -f docker/Dockerfile -t myapp:latest .

# Run with Docker Compose
docker-run:
	docker-compose -f docker/docker-compose.yml up --build

docker-down:
	docker-compose -f docker/docker-compose.yml down

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Development server with hot reload (requires air)
dev:
	air

# psql command
psql:
	docker exec -it simple-bank-db psql -U chinsiang -d go-project

# migrations
migrate-up:
	migrate -path internal/database/migrations -database "postgresql://chinsiang:chinsiang@localhost:5432/go-project?sslmode=disable" -verbose up

migrate-down:
	migrate -path internal/database/migrations -database "postgresql://chinsiang:chinsiang@localhost:5432/go-project?sslmode=disable" -verbose down

# sqlc
sqlc:
	sqlc generate
