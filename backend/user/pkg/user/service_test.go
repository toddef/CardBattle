package user

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository is a mock implementation of Repository
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateUser(ctx context.Context, user *User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockRepository) GetUserByID(ctx context.Context, id string) (*User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockRepository) UpdateUser(ctx context.Context, id string, updates *UpdateUserRequest) error {
	args := m.Called(ctx, id, updates)
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	testCases := []struct {
		name    string
		req     *CreateUserRequest
		mockErr error
		wantErr bool
	}{
		{
			name: "successful creation",
			req: &CreateUserRequest{
				ID:       "123",
				Username: "testuser",
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name: "missing required fields",
			req: &CreateUserRequest{
				ID: "123",
			},
			mockErr: nil,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.wantErr {
				mockRepo.On("CreateUser", ctx, mock.AnythingOfType("*user.User")).Return(tc.mockErr)
			}

			user, err := service.CreateUser(ctx, tc.req)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.req.ID, user.ID)
			assert.Equal(t, tc.req.Username, user.Username)
		})
	}
}

func TestGetUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	testUser := &User{
		ID:        "123",
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	testCases := []struct {
		name    string
		id      string
		mock    bool
		mockErr error
		want    *User
		wantErr bool
	}{
		{
			name:    "user found",
			id:      "123",
			mock:    true,
			mockErr: nil,
			want:    testUser,
			wantErr: false,
		},
		{
			name:    "empty id",
			id:      "",
			mock:    false,
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.mock {
				mockRepo.On("GetUserByID", ctx, tc.id).Return(tc.want, tc.mockErr)
			}

			user, err := service.GetUser(ctx, tc.id)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.want.ID, user.ID)
			assert.Equal(t, tc.want.Username, user.Username)
		})
	}
}

func TestUpdateUser(t *testing.T) {
	mockRepo := new(MockRepository)
	service := NewUserService(mockRepo)
	ctx := context.Background()

	testUser := &User{
		ID:        "123",
		Username:  "testuser",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	newUsername := "newuser"
	testCases := []struct {
		name    string
		id      string
		req     *UpdateUserRequest
		mockErr error
		wantErr bool
	}{
		{
			name: "successful update",
			id:   "123",
			req: &UpdateUserRequest{
				Username: &newUsername,
			},
			mockErr: nil,
			wantErr: false,
		},
		{
			name:    "empty id",
			id:      "",
			req:     &UpdateUserRequest{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if !tc.wantErr {
				mockRepo.On("GetUserByID", ctx, tc.id).Return(testUser, nil)
				mockRepo.On("UpdateUser", ctx, tc.id, tc.req).Return(tc.mockErr)
				mockRepo.On("GetUserByID", ctx, tc.id).Return(testUser, nil)
			}

			user, err := service.UpdateUser(ctx, tc.id, tc.req)

			if tc.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, testUser.ID, user.ID)
		})
	}
}
