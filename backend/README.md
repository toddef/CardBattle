# Card Battle Backend Services

This directory contains the microservices that power the Card Battle game.

## Services

- Auth Service (port 8080): Authentication and authorization
- User Service (port 8081): User management and profiles
- Deck Service (port 8082): Deck creation and management
- Card Service (port 8083): Card database and operations
- Lobby Service (port 8084): Game lobby and matchmaking
- Game Service (port 8085): Game state and mechanics
- Replay Service (port 8086): Game replay storage and playback

## Development

### Prerequisites

- Go 1.24 or higher
- Make (optional, for convenience)

### Running Locally

To run a service locally, navigate to its cmd directory and run:

```bash
go run main.go
```

For example, to run the Auth service:

```bash
cd auth/cmd
go run main.go
```

Each service will start on its designated port. You can verify it's running by checking the health endpoint:

```bash
curl http://localhost:808X/healthz
```

Replace X with the service port number (0-6).

### Running Tests

To run tests for all services:

```bash
# From the backend directory
go test ./... -v
```

To run tests for a specific service:

```bash
# From the backend directory
go test ./servicename/... -v
```

### Service Ports

- Auth: 8080
- User: 8081
- Deck: 8082
- Card: 8083
- Lobby: 8084
- Game: 8085
- Replay: 8086

Each service exposes a `/healthz` endpoint that returns a 200 OK response with `{"status": "ok"}` when the service is healthy. 