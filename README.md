# Card Battle Game Monorepo

This monorepo contains a card-battling game implementation split across multiple services and components.

## Project Structure

### Backend Services (`/backend`)
- `auth/` - Authentication and authorization service
- `card/` - Card management and metadata service
- `deck/` - Deck building and management service
- `game/` - Core game logic and battle mechanics
- `lobby/` - Game lobby and matchmaking service
- `replay/` - Game replay storage and playback service
- `user/` - User profile and management service
- `pkg/` - Shared libraries and utilities

### Frontend (`/frontend`)
- React-based web application
- Modern UI with responsive design
- Built with Vite and TypeScript
- Nginx-based container deployment

### Infrastructure (`/infrastructure`)
- `helm/` - Helm charts for Kubernetes deployment
  - `card-battle/` - Main Helm chart for all services
- `argocd/` - ArgoCD configuration for GitOps deployment

## Prerequisites

- Go 1.24 or later
- Node.js 20 or later
- Docker
- Kubernetes cluster (Minikube, Docker Desktop, or cloud-based)
- Helm
- kubectl

## Building the Project

### Backend Services

```bash
# From the root directory
cd backend
./build-services.ps1  # Windows
./build-services.sh   # Linux/macOS
```

This will build Docker images for all backend services.

### Frontend

```bash
# From the root directory
cd frontend
./build-frontend.ps1  # Windows
./build-frontend.sh   # Linux/macOS
```

### Combined Build

```bash
# From the root directory
./build-all.ps1  # Windows
./build-all.sh   # Linux/macOS
```

## Deployment

### Local Development with Minikube

1. Start Minikube:
```bash
minikube start
```

2. Enable the ingress addon:
```bash
minikube addons enable ingress
```

3. Point Docker to Minikube's daemon:
```bash
# Windows PowerShell
minikube docker-env | Invoke-Expression
# Linux/macOS
eval $(minikube docker-env)
```

4. Build all services:
```bash
./build-all.ps1  # Windows
./build-all.sh   # Linux/macOS
```

5. Deploy using Helm:
```bash
helm install card-battle ./infrastructure/helm/card-battle
```

### Production Deployment

The project is configured for GitOps deployment using ArgoCD. See `/infrastructure/argocd` for configuration details.

## Architecture

### Backend Services
- Microservices architecture using Go
- gRPC for inter-service communication
- REST APIs for frontend communication
- Shared libraries in `pkg/` for common functionality

### Frontend
- React-based SPA
- TypeScript for type safety
- Modern UI components
- Responsive design
- Nginx for static file serving

### Infrastructure
- Kubernetes-native deployment
- Helm for package management
- Ingress for routing
- Service mesh ready

## Development Workflow

1. Make changes to services
2. Build using the appropriate build script
3. Deploy using Helm
4. Access the application through the ingress

## Current Status

- Backend services: Implemented and containerized
- Frontend: Basic implementation with menu system
- Infrastructure: Helm charts and ArgoCD configuration ready
- Local development environment: Configured for Minikube
- CI/CD: GitHub Actions workflows (in progress)

## Next Steps

- Complete frontend implementation
- Add automated testing
- Set up monitoring and logging
- Implement WebSocket support for real-time gameplay
- Add documentation for API endpoints

## License

TBD
