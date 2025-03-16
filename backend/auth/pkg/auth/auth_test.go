package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockGoogleOAuthClient mocks the Google OAuth client
type MockGoogleOAuthClient struct {
	mock.Mock
}

func (m *MockGoogleOAuthClient) Exchange(code string) (*GoogleUserInfo, error) {
	args := m.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*GoogleUserInfo), args.Error(1)
}

func TestGoogleOAuthHandler(t *testing.T) {
	// Create mock client
	mockClient := new(MockGoogleOAuthClient)

	// Setup test cases
	testCases := []struct {
		name         string
		code         string
		mockResponse *GoogleUserInfo
		mockError    error
		expectedCode int
		expectCookie bool
	}{
		{
			name: "successful oauth flow",
			code: "valid_code",
			mockResponse: &GoogleUserInfo{
				ID:    "123",
				Email: "test@example.com",
				Name:  "Test User",
			},
			mockError:    nil,
			expectedCode: http.StatusOK,
			expectCookie: true,
		},
		{
			name:         "invalid code",
			code:         "invalid_code",
			mockResponse: nil,
			mockError:    ErrInvalidCode,
			expectedCode: http.StatusBadRequest,
			expectCookie: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup mock expectations
			if tc.mockError == nil {
				mockClient.On("Exchange", tc.code).Return(tc.mockResponse, nil)
			} else {
				mockClient.On("Exchange", tc.code).Return(nil, tc.mockError)
			}

			// Create auth service with mock client
			authService := NewAuthService(mockClient)

			// Create test request
			req := httptest.NewRequest("GET", "/oauth/google?code="+tc.code, nil)
			w := httptest.NewRecorder()

			// Handle request
			authService.HandleGoogleOAuth(w, req)

			// Assert response
			assert.Equal(t, tc.expectedCode, w.Code)

			if tc.expectCookie {
				cookies := w.Result().Cookies()
				assert.NotEmpty(t, cookies)

				// Find JWT cookie
				var jwtCookie *http.Cookie
				for _, cookie := range cookies {
					if cookie.Name == "jwt" {
						jwtCookie = cookie
						break
					}
				}

				assert.NotNil(t, jwtCookie)
				assert.True(t, jwtCookie.HttpOnly)
				assert.Equal(t, "/", jwtCookie.Path)

				// Validate JWT token
				token, err := authService.ValidateToken(jwtCookie.Value)
				assert.NoError(t, err)
				assert.Equal(t, tc.mockResponse.ID, token.Claims.(*Claims).UserID)
			}
		})
	}
}

func TestJWTGenerationAndValidation(t *testing.T) {
	authService := NewAuthService(nil)
	userID := "123"

	// Generate token
	token, err := authService.GenerateToken(userID)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Validate token
	claims, err := authService.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, userID, claims.Claims.(*Claims).UserID)
}

func TestCookieGeneration(t *testing.T) {
	authService := NewAuthService(nil)
	token := "test.jwt.token"

	cookie := authService.CreateJWTCookie(token)
	assert.Equal(t, "jwt", cookie.Name)
	assert.Equal(t, token, cookie.Value)
	assert.True(t, cookie.HttpOnly)
	assert.Equal(t, "/", cookie.Path)
	assert.True(t, cookie.Secure)
	assert.Equal(t, http.SameSiteStrictMode, cookie.SameSite)
}
