package handler

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/Mahathirrr/task-management-backend/internal/service"
	"github.com/Mahathirrr/task-management-backend/pkg/oauth"
	"github.com/Mahathirrr/task-management-backend/pkg/response"
)

type OAuthHandler struct {
	authService  service.AuthService
	oauthManager *oauth.OAuthManager
}

func NewOAuthHandler(authService service.AuthService, oauthManager *oauth.OAuthManager) *OAuthHandler {
	return &OAuthHandler{
		authService:  authService,
		oauthManager: oauthManager,
	}
}

// GoogleAuth redirects to Google OAuth
func (h *OAuthHandler) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	h.handleOAuthRedirect(w, r, "google")
}

// GoogleCallback handles Google OAuth callback
func (h *OAuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
	h.handleOAuthCallback(w, r, "google")
}

// handleOAuthRedirect handles OAuth redirect for any provider
func (h *OAuthHandler) handleOAuthRedirect(w http.ResponseWriter, r *http.Request, providerName string) {
	provider, err := h.oauthManager.GetProvider(providerName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "OAuth provider not supported")
		return
	}

	// Generate secure state parameter
	state, err := generateSecureState()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to generate state")
		return
	}

	// Store state in session/cookie for validation
	// Use more permissive cookie settings for development
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HttpOnly: true,
		Secure:   false,                            // Set to false for development (HTTP)
		SameSite: http.SameSiteLaxMode,             // Changed to Lax for better compatibility
		Expires:  time.Now().Add(10 * time.Minute), // Short-lived
		Path:     "/",                              // Changed to root path for better accessibility
	})

	// Get authorization URL
	authURL, err := provider.GetAuthURL(state)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get authorization URL")
		return
	}

	// Redirect to OAuth provider
	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// handleOAuthCallback handles OAuth callback for any provider
func (h *OAuthHandler) handleOAuthCallback(w http.ResponseWriter, r *http.Request, providerName string) {
	// Get code and state from query parameters
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if code == "" {
		response.Error(w, http.StatusBadRequest, "Missing authorization code")
		return
	}

	if state == "" {
		response.Error(w, http.StatusBadRequest, "Missing state parameter")
		return
	}

	// Validate state parameter
	stateCookie, err := r.Cookie("oauth_state")
	if err != nil {
		// Log the error for debugging
		response.Error(w, http.StatusBadRequest, "State cookie not found. Please try logging in again.")
		return
	}

	if stateCookie.Value != state {
		response.Error(w, http.StatusBadRequest, "Invalid state parameter. Please try logging in again.")
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		HttpOnly: true,
		Secure:   false, // Match the original cookie settings
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/",
	})

	// Get OAuth provider
	provider, err := h.oauthManager.GetProvider(providerName)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "OAuth provider not supported")
		return
	}

	// Get user info from OAuth provider
	oauthUser, err := provider.GetUserInfo(code, state)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to get user info from OAuth provider")
		return
	}

	// Process OAuth login
	authResp, err := h.authService.OAuthLogin(oauthUser)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to process OAuth login")
		return
	}

	// Set refresh token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    authResp.RefreshToken,
		HttpOnly: true,
		Secure:   false, // Set to false for development
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(168 * time.Hour), // 7 days
		Path:     "/api/v1/auth",
	})

	// For development, we can redirect to frontend with token in URL
	// In production, you might want to use a different approach
	frontendURL := "http://localhost:3000/auth/callback"
	redirectURL := frontendURL + "?token=" + authResp.AccessToken + "&success=true"

	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

// generateSecureState generates a secure random state parameter
func generateSecureState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
