package user

import (
	"context"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

// PostgresRepository implements Repository interface using PostgreSQL
type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *PostgresRepository) CreateUser(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (id, username, avatar_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, username, avatar_url, created_at, updated_at`

	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := r.db.QueryRowxContext(
		ctx,
		query,
		user.ID,
		user.Username,
		user.AvatarURL,
		user.CreatedAt,
		user.UpdatedAt,
	).StructScan(user)

	return err
}

// GetUserByID retrieves a user by their ID
func (r *PostgresRepository) GetUserByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, username, avatar_url, created_at, updated_at
		FROM users
		WHERE id = $1`

	var user User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates user information
func (r *PostgresRepository) UpdateUser(ctx context.Context, id string, updates *UpdateUserRequest) error {
	// Build dynamic query based on provided fields
	query := `UPDATE users SET updated_at = NOW()`
	args := []interface{}{id}
	argPosition := 2

	if updates.Username != nil {
		query += `, username = $` + string(rune('0'+argPosition))
		args = append(args, *updates.Username)
		argPosition++
	}

	if updates.AvatarURL != nil {
		query += `, avatar_url = $` + string(rune('0'+argPosition))
		args = append(args, *updates.AvatarURL)
		argPosition++
	}

	query += ` WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("user not found")
	}

	return nil
}
