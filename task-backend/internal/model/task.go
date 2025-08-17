package model

import "time"

type Task struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
)

// TaskCreateRequest for creating task
type TaskCreateRequest struct {
	Title       string     `json:"title" validate:"required,max=255"`
	Description *string    `json:"description"`
	Status      TaskStatus `json:"status" validate:"omitempty,oneof=pending in_progress completed"`
}

// TaskUpdateRequest for updating task
type TaskUpdateRequest struct {
	Title       *string     `json:"title,omitempty" validate:"omitempty,max=255"`
	Description *string     `json:"description,omitempty"`
	Status      *TaskStatus `json:"status,omitempty" validate:"omitempty,oneof=pending in_progress completed"`
}

// Response DTOs

// TasksResponse for paginated tasks response
type TasksResponse struct {
	Tasks []Task `json:"tasks"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
