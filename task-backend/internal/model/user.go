package model

import "time"

type User struct {
	ID            int        `json:"id"`
	Email         string     `json:"email"`
	Name          string     `json:"name"`
	Password      string     `json:"-"` // Hide password from JSON
	Role          UserRole   `json:"role,omitempty"`
	OauthProvider *string    `json:"oauth_provider,omitempty"`
	OauthID       *string    `json:"oauth_id,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}

type UserRole string

const (
	UserRoleUser  UserRole = "user"
	UserRoleAdmin UserRole = "admin"
)

type UserRegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"-"` // Hidden from JSON response
	User        User   `json:"user"`
}

// UsersResponse for paginated users response
type UsersResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
