# Service definitions with their ports
$services = @{
    "card" = 8081
    "deck" = 8082
    "game" = 8083
    "lobby" = 8084
    "replay" = 8085
}

# Read the template
$template = Get-Content -Path "Dockerfile.template" -Raw

# Generate Dockerfile for each service
foreach ($service in $services.GetEnumerator()) {
    $serviceName = $service.Key
    $servicePort = $service.Value
    
    Write-Host "Generating Dockerfile for $serviceName service..."
    
    # Replace placeholders
    $dockerfile = $template.Replace('${SERVICE_NAME}', $serviceName).Replace('${SERVICE_PORT}', $servicePort)
    
    # Write to service directory
    $dockerfile | Set-Content -Path "$serviceName/Dockerfile"
}

Write-Host "All Dockerfiles generated successfully!" 