# Build stage
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Copy backend shared code first
COPY pkg /app/pkg
COPY go.mod go.sum /app/

# Copy user service code
COPY user /app/user/

# Set working directory to user service
WORKDIR /app/user

# Download dependencies
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /user-service ./cmd

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /user-service .

# Expose port
EXPOSE 8086

# Run the service
CMD ["./user-service"] 
