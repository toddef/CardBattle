#!/bin/bash

# Service definitions with their ports
declare -A services
services=(
    ["auth"]=8080
    ["card"]=8081
    ["deck"]=8082
    ["game"]=8083
    ["lobby"]=8084
    ["replay"]=8085
    ["user"]=8086
)

# Read the template
template=$(cat Dockerfile.template)

# Generate Dockerfile for each service
for service in "${!services[@]}"; do
    port="${services[$service]}"
    echo "Generating Dockerfile for $service service..."
    
    # Replace placeholders
    dockerfile="${template//\${SERVICE_NAME\}/$service}"
    dockerfile="${dockerfile//\${SERVICE_PORT\}/$port}"
    
    # Write to service directory
    echo "$dockerfile" > "$service/Dockerfile"
done

echo "All Dockerfiles generated successfully!" 