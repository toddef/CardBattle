#!/bin/bash

# List of services
services=("deck" "game" "lobby" "replay" "user")

# Base Dockerfile content
dockerfile_content='# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

# Pre-copy/cache go.mod for pre-downloading dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /$SERVICE-service ./$SERVICE/cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /$SERVICE-service .

# Expose port (will be specified in Kubernetes)
EXPOSE 8080

# Run the service
CMD ["./$SERVICE-service"]'

# Generate Dockerfile for each service
for service in "${services[@]}"; do
  echo "${dockerfile_content/\$SERVICE/$service}" > "$service/Dockerfile"
  echo "Created Dockerfile for $service service"
done 