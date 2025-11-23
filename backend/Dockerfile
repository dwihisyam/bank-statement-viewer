# ========== 1. Build Stage ==========
FROM golang:1.22-alpine AS builder

# Install git (required for go mod download if using private repos)
RUN apk add --no-cache git

# Set working directory inside container
WORKDIR /app

# Copy go mod files first (layer caching)
COPY go.mod ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go binary
RUN go build -o server ./cmd/app

# ========== 2. Runtime Stage ==========
FROM alpine:3.20

# Create non-root user
RUN adduser -D appuser

# Set working directory inside container
WORKDIR /app

# Copy compiled binary from build stage
COPY --from=builder /app/server .

# Expose backend port
EXPOSE 8080

# Change user
USER appuser

# Run the binary
CMD ["./server"]
