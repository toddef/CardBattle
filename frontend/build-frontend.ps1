# Configuration
$registry = if ($env:DOCKER_REGISTRY) { $env:DOCKER_REGISTRY } else { "card-battle" }
$tag = if ($env:BUILD_TAG) { $env:BUILD_TAG } else { "latest" }

Write-Host "Building frontend..."
Write-Host "Using registry: $registry"
Write-Host "Using tag: $tag"

# Build the frontend container
docker build -t "$registry/frontend:$tag" .

# If we're in GitHub Actions, also push the image
if ($env:GITHUB_ACTIONS) {
    Write-Host "Pushing frontend image to registry..."
    docker push "$registry/frontend:$tag"
}

Write-Host "Frontend build completed successfully!" 