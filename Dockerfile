# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

# Install git (needed for Go modules)
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary
RUN go build -o pismoapp main.go

# Stage 2: Create a minimal image
FROM debian:bookworm-slim

# Set working directory
WORKDIR /root/

# Copy the compiled binary from builder
COPY --from=builder /app/pismoapp .

# Expose port your API listens on
EXPOSE 8080

# Command to run the binary
CMD ["./pismoapp"]