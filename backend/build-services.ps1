# List of services
$services = @{
    "auth" = 8080
    "card" = 8081
    "deck" = 8082
    "game" = 8083
    "lobby" = 8084
    "replay" = 8085
    "user" = 8086
}

# Generate Dockerfiles first
.\generate-dockerfiles.ps1

# Build all services
foreach ($service in $services.GetEnumerator()) {
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
    docker build -t "card-battle/$serviceName-service:latest" .build-context
    
    # Clean up
    Remove-Item -Path ".build-context" -Recurse -Force
    
    # If we're in GitHub Actions, also push the image
    if ($env:GITHUB_ACTIONS) {
        Write-Host "Pushing $serviceName service to registry..."
        docker push "card-battle/$serviceName-service:latest"
    }
}

Write-Host "All services built successfully!" 