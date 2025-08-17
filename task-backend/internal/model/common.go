package model

// ErrorResponse untuk error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// ValidationError untuk validation error details
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrorResponse untuk validation error response
type ValidationErrorResponse struct {
	Error   string            `json:"error"`
	Details []ValidationError `json:"details"`
}

// SuccessResponse untuk generic success response
type SuccessResponse struct {
	Message string `json:"message"`
}

// TokenResponse untuk refresh token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

// Common constants untuk response messages
const (
	ErrInvalidCredentials = "Invalid credentials"
	ErrEmailAlreadyExists = "Email already exists"
	ErrTaskNotFound       = "Task not found"
	ErrUserNotFound       = "User not found"
	ErrUnauthorized       = "Unauthorized"
	ErrForbidden          = "Access denied"
	ErrValidationFailed   = "Validation failed"
	ErrInternalServer     = "Internal server error"

	MsgLoginSuccess    = "Login successful"
	MsgLogoutSuccess   = "Logout successful"
	MsgRegisterSuccess = "Registration successful"
	MsgTaskCreated     = "Task created successfully"
	MsgTaskUpdated     = "Task updated successfully"
	MsgTaskDeleted     = "Task deleted successfully"
	MsgUserDeleted     = "User deleted successfully"
)
