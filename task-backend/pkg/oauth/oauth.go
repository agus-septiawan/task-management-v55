package oauth

import (
	"fmt"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

// OAuthProvider interface untuk semua OAuth providers
type OAuthProvider interface {
	GetAuthURL(state string) (string, error)
	GetUserInfo(code, state string) (*OAuthUser, error)
}

// OAuthUser represents user data from OAuth provider
type OAuthUser struct {
	ID       string
	Email    string
	Name     string
	Provider string
}

// OAuthManager manages OAuth providers
type OAuthManager struct {
	providers map[string]OAuthProvider
}

// NewOAuthManager creates new OAuth manager
func NewOAuthManager() *OAuthManager {
	return &OAuthManager{
		providers: make(map[string]OAuthProvider),
	}
}

// RegisterProvider registers new OAuth provider
func (m *OAuthManager) RegisterProvider(name string, provider OAuthProvider) {
	m.providers[name] = provider
}

// GetProvider returns OAuth provider by name
func (m *OAuthManager) GetProvider(name string) (OAuthProvider, error) {
	provider, exists := m.providers[name]
	if !exists {
		return nil, fmt.Errorf("OAuth provider %s not found", name)
	}
	return provider, nil
}

// GoogleProvider implements OAuth for Google
type GoogleProvider struct {
	clientID     string
	clientSecret string
	redirectURL  string
}

// NewGoogleProvider creates new Google OAuth provider
func NewGoogleProvider(clientID, clientSecret, redirectURL string) *GoogleProvider {
	// Initialize Goth Google provider
	goth.UseProviders(
		google.New(clientID, clientSecret, redirectURL),
	)
	return &GoogleProvider{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
	}
}

// GetAuthURL returns Google OAuth authorization URL
func (g *GoogleProvider) GetAuthURL(state string) (string, error) {
	provider, err := goth.GetProvider("google")
	if err != nil {
		return "", fmt.Errorf("failed to get Google provider: %w", err)
	}
	session, err := provider.BeginAuth(state)
	if err != nil {
		return "", fmt.Errorf("failed to begin auth: %w", err)
	}
	url, err := session.GetAuthURL()
	if err != nil {
		return "", fmt.Errorf("failed to get auth URL: %w", err)
	}
	return url, nil
}

// CustomParams implements goth.Params
type CustomParams map[string]string

// Get implements the Get method for goth.Params
func (c CustomParams) Get(key string) string {
	return c[key]
}

// GetUserInfo gets user information from Google OAuth
func (g *GoogleProvider) GetUserInfo(code, state string) (*OAuthUser, error) {
	provider, err := goth.GetProvider("google")
	if err != nil {
		return nil, fmt.Errorf("failed to get Google provider: %w", err)
	}
	session, err := provider.BeginAuth(state)
	if err != nil {
		return nil, fmt.Errorf("failed to begin auth: %w", err)
	}

	// Convert map to CustomParams
	params := CustomParams{
		"code": code,
	}

	// Complete the auth process
	_, err = session.Authorize(provider, params)
	if err != nil {
		return nil, fmt.Errorf("failed to authorize: %w", err)
	}

	user, err := provider.FetchUser(session)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &OAuthUser{
		ID:       user.UserID,
		Email:    user.Email,
		Name:     user.Name,
		Provider: "google",
	}, nil
}

// InitializeOAuth initializes OAuth providers
func InitializeOAuth(googleClientID, googleClientSecret, googleRedirectURL string) *OAuthManager {
	manager := NewOAuthManager()
	// Register Google provider
	googleProvider := NewGoogleProvider(googleClientID, googleClientSecret, googleRedirectURL)
	manager.RegisterProvider("google", googleProvider)
	return manager
}
