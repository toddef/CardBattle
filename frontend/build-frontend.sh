#!/bin/bash

# Configuration
REGISTRY=${DOCKER_REGISTRY:-"card-battle"}
TAG=${BUILD_TAG:-"latest"}

echo "Building frontend..."
echo "Using registry: $REGISTRY"
echo "Using tag: $TAG"

# Build the frontend container
docker build -t "$REGISTRY/frontend:$TAG" .

# If we're in GitHub Actions, also push the image
if [ ! -z "$GITHUB_ACTIONS" ]; then
    echo "Pushing frontend image to registry..."
    docker push "$REGISTRY/frontend:$TAG"
fi

echo "Frontend build completed successfully!" 