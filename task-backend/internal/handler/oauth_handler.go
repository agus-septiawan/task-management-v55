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
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    state,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(10 * time.Minute), // Short-lived
		Path:     "/api/v1/auth/oauth",
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

	if code == "" || state == "" {
		response.Error(w, http.StatusBadRequest, "Missing code or state parameter")
		return
	}

	// Validate state parameter
	stateCookie, err := r.Cookie("oauth_state")
	if err != nil || stateCookie.Value != state {
		response.Error(w, http.StatusBadRequest, "Invalid state parameter")
		return
	}

	// Clear state cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "oauth_state",
		Value:    "",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(-time.Hour),
		Path:     "/api/v1/auth/oauth",
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
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(168 * time.Hour), // 7 days
		Path:     "/api/v1/auth",
	})

	response.JSON(w, http.StatusOK, authResp)
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
