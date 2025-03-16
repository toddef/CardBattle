package auth

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrInvalidCode = errors.New("invalid oauth code")
	jwtSecret      = []byte(os.Getenv("JWT_SECRET"))
)

// GoogleUserInfo represents the user information returned by Google
type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
}

// GoogleOAuthClient interface for dependency injection and testing
type GoogleOAuthClient interface {
	Exchange(code string) (*GoogleUserInfo, error)
}

// Claims represents the JWT claims structure
type Claims struct {
	jwt.RegisteredClaims
	UserID string `json:"user_id"`
}

// AuthService handles authentication operations
type AuthService struct {
	googleClient GoogleOAuthClient
}

// NewAuthService creates a new AuthService instance
func NewAuthService(googleClient GoogleOAuthClient) *AuthService {
	return &AuthService{
		googleClient: googleClient,
	}
}

// HandleGoogleOAuth handles the Google OAuth callback
func (s *AuthService) HandleGoogleOAuth(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "missing code parameter", http.StatusBadRequest)
		return
	}

	// Exchange code for user info
	userInfo, err := s.googleClient.Exchange(code)
	if err != nil {
		if err == ErrInvalidCode {
			http.Error(w, "invalid code", http.StatusBadRequest)
		} else {
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	// Generate JWT
	token, err := s.GenerateToken(userInfo.ID)
	if err != nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	// Set JWT cookie
	http.SetCookie(w, s.CreateJWTCookie(token))
	w.WriteHeader(http.StatusOK)
}

// GenerateToken creates a new JWT token for the given user ID
func (s *AuthService) GenerateToken(userID string) (string, error) {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates and parses a JWT token
func (s *AuthService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// CreateJWTCookie creates an HTTP cookie containing the JWT token
func (s *AuthService) CreateJWTCookie(token string) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(24 * time.Hour.Seconds()),
	}
}
