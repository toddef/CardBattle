# List of services
$services = @("auth", "card", "deck", "game", "lobby", "replay", "user")
$registry = if ($env:DOCKER_REGISTRY) { $env:DOCKER_REGISTRY } else { "card-battle" }

# Build backend services
foreach ($service in $services) {
    Write-Host "Building $service service..."
    docker build -t "$registry/$service-service:latest" -f "$service/Dockerfile" .
    
    # If we're in GitHub Actions, also push the image
    if ($env:GITHUB_ACTIONS) {
        Write-Host "Pushing $service service to registry..."
        docker push "$registry/$service-service:latest"
    }
}

# Build frontend
Write-Host "Building frontend..."
docker build -t "$registry/frontend:latest" -f "../frontend/Dockerfile" "../frontend"

# If we're in GitHub Actions, also push the frontend image
if ($env:GITHUB_ACTIONS) {
    Write-Host "Pushing frontend to registry..."
    docker push "$registry/frontend:latest"
}

Write-Host "All services built successfully!" 