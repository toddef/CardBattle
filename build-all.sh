#!/bin/bash

# Configuration
REGISTRY=${DOCKER_REGISTRY:-"card-battle"}
TAG=${BUILD_TAG:-"latest"}

echo "Building Card Battle application..."
echo "Using registry: $REGISTRY"
echo "Using tag: $TAG"

# Build backend services
echo -e "\nBuilding backend services..."
pushd backend
./build-services.sh
popd

# Build frontend
echo -e "\nBuilding frontend..."
pushd frontend
./build-frontend.sh
popd

echo -e "\nAll components built successfully!"

# If we're in GitHub Actions, list all images
if [ ! -z "$GITHUB_ACTIONS" ]; then
    echo -e "\nBuilt images:"
    docker images | grep "$REGISTRY"
fi 