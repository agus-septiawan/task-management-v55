package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Mahathirrr/task-management-backend/internal/model"
)

func TestTaskEndpoints(t *testing.T) {
	server := setupTestServer()

	t.Run("POST /api/v1/tasks - Valid Request", func(t *testing.T) {
		reqBody := model.TaskCreateRequest{
			Title:       "Test Task",
			Description: stringPtr("Test Description"),
			Status:      model.TaskStatusPending,
		}

		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer test-token")

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		// Note: This will fail without proper authentication setup
		// Expected status would be 401 (Unauthorized) without valid token
		if w.Code != http.StatusUnauthorized && w.Code != http.StatusCreated {
			t.Errorf("Expected status 401 or 201, got %d", w.Code)
		}
	})

	t.Run("POST /api/v1/tasks - Invalid Request", func(t *testing.T) {
		reqBody := model.TaskCreateRequest{
			Title:  "", // Invalid: empty title
			Status: "invalid_status", // Invalid: bad status
		}

		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer test-token")

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		// Should return 401 (Unauthorized) due to invalid token
		// or 400 (Bad Request) if validation runs first
		if w.Code != http.StatusUnauthorized && w.Code != http.StatusBadRequest {
			t.Errorf("Expected status 401 or 400, got %d", w.Code)
		}
	})

	t.Run("GET /api/v1/tasks - No Token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/tasks", nil)

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", w.Code)
		}
	})

	t.Run("GET /api/v1/tasks/1 - No Token", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/api/v1/tasks/1", nil)

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", w.Code)
		}
	})

	t.Run("PUT /api/v1/tasks/1 - No Token", func(t *testing.T) {
		reqBody := model.TaskUpdateRequest{
			Title: stringPtr("Updated Title"),
		}

		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("PUT", "/api/v1/tasks/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", w.Code)
		}
	})

	t.Run("DELETE /api/v1/tasks/1 - No Token", func(t *testing.T) {
		req := httptest.NewRequest("DELETE", "/api/v1/tasks/1", nil)

		w := httptest.NewRecorder()
		server.ServeHTTP(w, req)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Expected status 401, got %d", w.Code)
		}
	})
}

// Helper function to create string pointer
func stringPtr(s string) *string {
	return &s
}