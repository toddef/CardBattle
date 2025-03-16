# Configuration
$registry = if ($env:DOCKER_REGISTRY) { $env:DOCKER_REGISTRY } else { "card-battle" }
$tag = if ($env:BUILD_TAG) { $env:BUILD_TAG } else { "latest" }

Write-Host "Building Card Battle application..."
Write-Host "Using registry: $registry"
Write-Host "Using tag: $tag"

# Build backend services
Write-Host "`nBuilding backend services..."
Push-Location backend
.\build-services.ps1
Pop-Location

# Build frontend
Write-Host "`nBuilding frontend..."
Push-Location frontend
.\build-frontend.ps1
Pop-Location

Write-Host "`nAll components built successfully!"

# If we're in GitHub Actions, list all images
if ($env:GITHUB_ACTIONS) {
    Write-Host "`nBuilt images:"
    docker images | Select-String "$registry"
} 