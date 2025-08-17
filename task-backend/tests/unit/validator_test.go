package unit

import (
	"testing"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/pkg/validator"
)

func TestValidateTaskStruct(t *testing.T) {
	t.Run("ValidTaskCreateRequest", func(t *testing.T) {
		req := model.TaskCreateRequest{
			Title:  "Test Task",
			Status: model.TaskStatusPending,
		}

		errors := validator.ValidateStruct(req)
		if len(errors) != 0 {
			t.Errorf("Expected no validation errors, got %d errors", len(errors))
		}
	})

	t.Run("InvalidTaskCreateRequest", func(t *testing.T) {
		req := model.TaskCreateRequest{
			Title:  "", // Required field empty
			Status: "invalid_status", // Invalid status
		}

		errors := validator.ValidateStruct(req)
		if len(errors) == 0 {
			t.Error("Expected validation errors, got none")
		}

		// Should have errors for Title and Status
		expectedFields := map[string]bool{
			"Title":  false,
			"Status": false,
		}

		for _, err := range errors {
			if _, exists := expectedFields[err.Field]; exists {
				expectedFields[err.Field] = true
			}
		}

		for field, found := range expectedFields {
			if !found {
				t.Errorf("Expected validation error for field %s", field)
			}
		}
	})
}

func TestValidateStruct(t *testing.T) {
	t.Run("ValidUserRegisterRequest", func(t *testing.T) {
		req := model.UserRegisterRequest{
			Name:     "John Doe",
			Email:    "john@example.com",
			Password: "password123",
		}

		errors := validator.ValidateStruct(req)
		if len(errors) != 0 {
			t.Errorf("Expected no validation errors, got %d errors", len(errors))
		}
	})

	t.Run("InvalidUserRegisterRequest", func(t *testing.T) {
		req := model.UserRegisterRequest{
			Name:     "", // Required field empty
			Email:    "invalid-email", // Invalid email format
			Password: "123", // Too short
		}

		errors := validator.ValidateStruct(req)
		if len(errors) == 0 {
			t.Error("Expected validation errors, got none")
		}

		// Check specific errors
		expectedErrors := map[string]bool{
			"Name":     false,
			"Email":    false,
			"Password": false,
		}

		for _, err := range errors {
			if _, exists := expectedErrors[err.Field]; exists {
				expectedErrors[err.Field] = true
			}
		}

		for field, found := range expectedErrors {
			if !found {
				t.Errorf("Expected validation error for field %s", field)
			}
		}
	})

	t.Run("ValidUserLoginRequest", func(t *testing.T) {
		req := model.UserLoginRequest{
			Email:    "john@example.com",
			Password: "password123",
		}

		errors := validator.ValidateStruct(req)
		if len(errors) != 0 {
			t.Errorf("Expected no validation errors, got %d errors", len(errors))
		}
	})

	t.Run("InvalidUserLoginRequest", func(t *testing.T) {
		req := model.UserLoginRequest{
			Email:    "", // Required field empty
			Password: "", // Required field empty
		}

		errors := validator.ValidateStruct(req)
		if len(errors) == 0 {
			t.Error("Expected validation errors, got none")
		}

		// Should have errors for both Email and Password
		if len(errors) < 2 {
			t.Errorf("Expected at least 2 validation errors, got %d", len(errors))
		}
	})
}