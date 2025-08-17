package unit

import (
	"testing"

	"github.com/Mahathirrr/task-management-backend/internal/model"
)

// Mock TaskRepository for testing
type mockTaskRepository struct {
	tasks  map[int]*model.Task
	nextID int
}

func newMockTaskRepository() *mockTaskRepository {
	return &mockTaskRepository{
		tasks:  make(map[int]*model.Task),
		nextID: 1,
	}
}

func (m *mockTaskRepository) Create(task *model.Task) error {
	task.ID = m.nextID
	m.nextID++
	m.tasks[task.ID] = task
	return nil
}

func (m *mockTaskRepository) GetByID(id int) (*model.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return nil, nil
	}
	return task, nil
}

func (m *mockTaskRepository) GetByUserID(userID int, page, limit int, status, search string) ([]model.Task, int, error) {
	var tasks []model.Task
	for _, task := range m.tasks {
		if task.UserID == userID {
			if status != "" && string(task.Status) != status {
				continue
			}
			tasks = append(tasks, *task)
		}
	}
	return tasks, len(tasks), nil
}

func (m *mockTaskRepository) GetAll(page, limit int, status, search string) ([]model.Task, int, error) {
	var tasks []model.Task
	for _, task := range m.tasks {
		if status != "" && string(task.Status) != status {
			continue
		}
		tasks = append(tasks, *task)
	}
	return tasks, len(tasks), nil
}

func (m *mockTaskRepository) Update(task *model.Task) error {
	m.tasks[task.ID] = task
	return nil
}

func (m *mockTaskRepository) Delete(id int) error {
	delete(m.tasks, id)
	return nil
}

func (m *mockTaskRepository) IsOwner(taskID, userID int) (bool, error) {
	task, exists := m.tasks[taskID]
	if !exists {
		return false, nil
	}
	return task.UserID == userID, nil
}

func TestTaskValidation(t *testing.T) {
	t.Run("ValidTaskCreateRequest", func(t *testing.T) {
		req := model.TaskCreateRequest{
			Title:  "Test Task",
			Status: model.TaskStatusPending,
		}

		if req.Title == "" {
			t.Error("Expected title to be set")
		}

		if req.Status != model.TaskStatusPending {
			t.Errorf("Expected status to be pending, got %s", req.Status)
		}
	})

	t.Run("TaskStatusConstants", func(t *testing.T) {
		if model.TaskStatusPending != "pending" {
			t.Errorf("Expected TaskStatusPending to be 'pending', got %s", model.TaskStatusPending)
		}

		if model.TaskStatusInProgress != "in_progress" {
			t.Errorf("Expected TaskStatusInProgress to be 'in_progress', got %s", model.TaskStatusInProgress)
		}

		if model.TaskStatusCompleted != "completed" {
			t.Errorf("Expected TaskStatusCompleted to be 'completed', got %s", model.TaskStatusCompleted)
		}
	})
}