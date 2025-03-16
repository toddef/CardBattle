# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /build

# Install build dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the shared packages
COPY pkg/ ./pkg/

# Copy the service-specific code
COPY deck/ ./deck/

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /deck-service ./deck/cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /deck-service .

# Expose port
EXPOSE 8082

# Run the service
CMD ["./deck-service"] 
