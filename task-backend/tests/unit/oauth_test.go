package unit

import (
	"testing"

	"github.com/Mahathirrr/task-management-backend/pkg/oauth"
)

func TestOAuthManager(t *testing.T) {
	t.Run("NewOAuthManager", func(t *testing.T) {
		manager := oauth.NewOAuthManager()
		if manager == nil {
			t.Error("Expected OAuth manager to be created")
		}
	})

	t.Run("RegisterAndGetProvider", func(t *testing.T) {
		manager := oauth.NewOAuthManager()
		
		// Create a mock provider
		googleProvider := oauth.NewGoogleProvider("test-client-id", "test-client-secret", "http://localhost:8080/callback")
		
		// Register provider
		manager.RegisterProvider("google", googleProvider)
		
		// Get provider
		provider, err := manager.GetProvider("google")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if provider == nil {
			t.Error("Expected provider to be returned")
		}
	})

	t.Run("GetNonExistentProvider", func(t *testing.T) {
		manager := oauth.NewOAuthManager()
		
		_, err := manager.GetProvider("nonexistent")
		if err == nil {
			t.Error("Expected error for non-existent provider")
		}
	})

	t.Run("InitializeOAuth", func(t *testing.T) {
		manager := oauth.InitializeOAuth("test-client-id", "test-client-secret", "http://localhost:8080/callback")
		if manager == nil {
			t.Error("Expected OAuth manager to be initialized")
		}

		// Test that Google provider is registered
		_, err := manager.GetProvider("google")
		if err != nil {
			t.Errorf("Expected Google provider to be registered, got error: %v", err)
		}
	})
}

func TestGoogleProvider(t *testing.T) {
	t.Run("NewGoogleProvider", func(t *testing.T) {
		provider := oauth.NewGoogleProvider("test-client-id", "test-client-secret", "http://localhost:8080/callback")
		if provider == nil {
			t.Error("Expected Google provider to be created")
		}
	})

	// Note: Testing GetAuthURL and GetUserInfo would require mocking Goth
	// or setting up integration tests with actual OAuth flow
}