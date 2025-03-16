# List of services with shared module
$sharedServices = @{
    "card" = 8081
    "deck" = 8082
    "game" = 8083
    "lobby" = 8084
    "replay" = 8085
}

# Set registry from environment variable or use default
$registry = if ($env:DOCKER_REGISTRY) { $env:DOCKER_REGISTRY } else { "card-battle" }

# Generate Dockerfiles for shared services
.\generate-dockerfiles.ps1

# Build user service first (independent module)
Write-Host "Building user service..."
docker build -t "$registry/user-service:latest" ./user

# If we're in GitHub Actions, push the user service image
if ($env:GITHUB_ACTIONS) {
    Write-Host "Pushing user service to registry..."
    docker push "$registry/user-service:latest"
}

# Build all shared services
foreach ($service in $sharedServices.GetEnumerator()) {
    $serviceName = $service.Key
    Write-Host "Building $serviceName service..."
    
    # Create a temporary build context
    New-Item -ItemType Directory -Force -Path ".build-context" | Out-Null
    
    # Copy necessary files to build context
    Copy-Item go.mod, go.sum -Destination ".build-context"
    Copy-Item -Path "pkg" -Destination ".build-context" -Recurse
    Copy-Item -Path $serviceName -Destination ".build-context" -Recurse
    Copy-Item "$serviceName/Dockerfile" -Destination ".build-context"
    
    # Build from the temporary context
    docker build -t "$registry/$serviceName-service:latest" .build-context
    
    # Clean up
    Remove-Item -Path ".build-context" -Recurse -Force
    
    # If we're in GitHub Actions, also push the image
    if ($env:GITHUB_ACTIONS) {
        Write-Host "Pushing $serviceName service to registry..."
        docker push "$registry/$serviceName-service:latest"
    }
}

# Build frontend
Write-Host "Building frontend..."
docker build -t "$registry/frontend:latest" -f ../frontend/Dockerfile ../frontend

# If we're in GitHub Actions, also push the frontend image
if ($env:GITHUB_ACTIONS) {
    Write-Host "Pushing frontend to registry..."
    docker push "$registry/frontend:latest"
}

Write-Host "All services built successfully!" 