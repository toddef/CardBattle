# User Service

The User Service manages user profiles in the Card Battle game, storing data in PostgreSQL and integrating with the Auth Service.

## Prerequisites

- Go 1.24.1 or later
- PostgreSQL 14 or later
- golang-migrate CLI tool

## Setup

1. Install golang-migrate:
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

2. Set up environment variables:
```bash
export DATABASE_URL="postgres://username:password@localhost:5432/cardbattle?sslmode=disable"
export PORT=8086  # optional, defaults to 8086
```

3. Run database migrations:
```bash
migrate -database "${DATABASE_URL}" -path migrations up
```

## Running Tests

Run all tests:
```bash
go test ./...
```

Run specific test:
```bash
go test ./pkg/user -run TestCreateUser
```

## API Endpoints

### Create User
```http
POST /users
Content-Type: application/json

{
    "id": "auth0|123",
    "username": "player1",
    "avatar_url": "https://example.com/avatar.jpg"
}
```

### Get User
```http
GET /users/{id}
```

### Update User
```http
PATCH /users/{id}
Content-Type: application/json

{
    "username": "newname",
    "avatar_url": "https://example.com/new-avatar.jpg"
}
```

## Integration with Auth Service

The User Service automatically creates a new user profile when a user first logs in through the Auth Service. The Auth Service provides the user ID, which is used as the primary key in the users table.

## Development

1. Start PostgreSQL:
```bash
docker run -d --name postgres \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_DB=cardbattle \
    -p 5432:5432 \
    postgres:14
```

2. Run migrations:
```bash
migrate -database "postgres://postgres:postgres@localhost:5432/cardbattle?sslmode=disable" -path migrations up
```

3. Start the service:
```bash
go run cmd/main.go
``` 