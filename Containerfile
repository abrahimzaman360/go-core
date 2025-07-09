# Start from the official Golang image for building
FROM golang:1.24.4-alpine AS builder

# Set necessary environment variables
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Create and set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal base image for production
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the built binary from builder
COPY --from=builder /app/main .

# Expose application port (change if needed)
EXPOSE 3000

# Run the binary
CMD ["./main"]