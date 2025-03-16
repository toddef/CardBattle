package user

import (
	"context"
	"errors"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidInput = errors.New("invalid input")
)

// UserService implements the Service interface
type UserService struct {
	repo Repository
}

// NewUserService creates a new user service
func NewUserService(repo Repository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(ctx context.Context, req *CreateUserRequest) (*User, error) {
	if req.ID == "" || req.Username == "" {
		return nil, ErrInvalidInput
	}

	user := &User{
		ID:        req.ID,
		Username:  req.Username,
		AvatarURL: req.AvatarURL,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (s *UserService) GetUser(ctx context.Context, id string) (*User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// UpdateUser updates user information
func (s *UserService) UpdateUser(ctx context.Context, id string, req *UpdateUserRequest) (*User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	// Check if user exists
	_, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, ErrUserNotFound
	}

	// Update user
	if err := s.repo.UpdateUser(ctx, id, req); err != nil {
		return nil, err
	}

	// Get updated user
	return s.repo.GetUserByID(ctx, id)
}
