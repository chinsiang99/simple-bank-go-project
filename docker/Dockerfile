# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/server

# Final stage (use minimal image, not golang)
FROM alpine:latest AS final

RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app

COPY --from=builder /app/main .

# # Copy Swagger docs
# COPY internal/docs /app/docs

# Copy the Swagger docs from the builder
COPY --from=builder /app/internal/docs /app/internal/docs

# Let logs go to stdout/stderr instead of container fs
EXPOSE 8080

CMD ["./main"]
