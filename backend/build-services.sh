#!/bin/bash

# List of services
SERVICES=("auth" "card" "deck" "game" "lobby" "replay" "user")
REGISTRY=${DOCKER_REGISTRY:-"card-battle"}  # Use environment variable or default

# Generate Dockerfiles first
./generate-dockerfiles.sh

# Build all services
for service in "${SERVICES[@]}"; do
    echo "Building $service service..."
    
    # Create a temporary build context
    mkdir -p .build-context
    
    # Copy necessary files to build context
    cp go.mod go.sum .build-context/
    cp -r pkg .build-context/
    cp -r $service .build-context/
    cp $service/Dockerfile .build-context/
    
    # Build from the temporary context
    docker build -t $REGISTRY/$service-service:latest .build-context
    
    # Clean up
    rm -rf .build-context
    
    # If we're in GitHub Actions, also push the image
    if [ ! -z "$GITHUB_ACTIONS" ]; then
        echo "Pushing $service service to registry..."
        docker push $REGISTRY/$service-service:latest
    fi
done

# Build frontend
echo "Building frontend..."
docker build -t $REGISTRY/frontend:latest -f ../frontend/Dockerfile ../frontend

# If we're in GitHub Actions, also push the frontend image
if [ ! -z "$GITHUB_ACTIONS" ]; then
    echo "Pushing frontend to registry..."
    docker push $REGISTRY/frontend:latest
fi

echo "All services built successfully!" 