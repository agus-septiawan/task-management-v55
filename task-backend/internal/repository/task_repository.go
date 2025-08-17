package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/Mahathirrr/task-management-backend/internal/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id int) (*model.Task, error)
	GetByUserID(userID int, page, limit int, status, search string) ([]model.Task, int, error)
	GetAll(page, limit int, status, search string) ([]model.Task, int, error)
	Update(task *model.Task) error
	Delete(id int) error
	IsOwner(taskID, userID int) (bool, error)
}

type taskRepository struct {
	db *sql.DB
}

// NewTaskRepository membuat instance TaskRepository
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

// Create membuat task baru
func (r *taskRepository) Create(task *model.Task) error {
	query := `
		INSERT INTO tasks (user_id, title, description, status)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(query, task.UserID, task.Title, task.Description, task.Status)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	task.ID = int(id)
	return nil
}

// GetByID mengambil task berdasarkan ID
func (r *taskRepository) GetByID(id int) (*model.Task, error) {
	query := `
		SELECT id, user_id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE id = ?
	`

	var task model.Task
	var description sql.NullString
	
	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get task by id: %w", err)
	}

	if description.Valid {
		task.Description = &description.String
	}

	return &task, nil
}

// GetByUserID mengambil tasks berdasarkan user ID dengan pagination dan filter
func (r *taskRepository) GetByUserID(userID int, page, limit int, status, search string) ([]model.Task, int, error) {
	offset := (page - 1) * limit
	
	// Build query with filters
	whereClause := "WHERE user_id = ?"
	args := []interface{}{userID}
	
	if status != "" {
		whereClause += " AND status = ?"
		args = append(args, status)
	}
	
	if search != "" {
		whereClause += " AND (title LIKE ? OR description LIKE ?)"
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, title, description, status, created_at, updated_at
		FROM tasks
		%s
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, whereClause)
	
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get tasks: %w", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		var description sql.NullString
		
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan task: %w", err)
		}
		
		if description.Valid {
			task.Description = &description.String
		}
		
		tasks = append(tasks, task)
	}

	// Count total tasks
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM tasks %s", whereClause)
	countArgs := args[:len(args)-2] // Remove LIMIT and OFFSET
	
	var total int
	err = r.db.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count tasks: %w", err)
	}

	return tasks, total, nil
}

// GetAll mengambil semua tasks dengan pagination dan filter (admin only)
func (r *taskRepository) GetAll(page, limit int, status, search string) ([]model.Task, int, error) {
	offset := (page - 1) * limit
	
	// Build query with filters
	var whereClause string
	var args []interface{}
	
	conditions := []string{}
	
	if status != "" {
		conditions = append(conditions, "status = ?")
		args = append(args, status)
	}
	
	if search != "" {
		conditions = append(conditions, "(title LIKE ? OR description LIKE ?)")
		searchPattern := "%" + search + "%"
		args = append(args, searchPattern, searchPattern)
	}
	
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT id, user_id, title, description, status, created_at, updated_at
		FROM tasks
		%s
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?
	`, whereClause)
	
	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get all tasks: %w", err)
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		var description sql.NullString
		
		err := rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan task: %w", err)
		}
		
		if description.Valid {
			task.Description = &description.String
		}
		
		tasks = append(tasks, task)
	}

	// Count total tasks
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM tasks %s", whereClause)
	countArgs := args[:len(args)-2] // Remove LIMIT and OFFSET
	
	var total int
	err = r.db.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count all tasks: %w", err)
	}

	return tasks, total, nil
}

// Update mengupdate task
func (r *taskRepository) Update(task *model.Task) error {
	query := `
		UPDATE tasks
		SET title = ?, description = ?, status = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.db.Exec(query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}

// Delete menghapus task
func (r *taskRepository) Delete(id int) error {
	query := "DELETE FROM tasks WHERE id = ?"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}

// IsOwner mengecek apakah user adalah pemilik task
func (r *taskRepository) IsOwner(taskID, userID int) (bool, error) {
	query := "SELECT user_id FROM tasks WHERE id = ?"
	
	var ownerID int
	err := r.db.QueryRow(query, taskID).Scan(&ownerID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, fmt.Errorf("failed to check task ownership: %w", err)
	}

	return ownerID == userID, nil
}