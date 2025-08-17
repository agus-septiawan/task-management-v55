package service

import (
	"errors"
	"fmt"

	"github.com/Mahathirrr/task-management-backend/internal/model"
	"github.com/Mahathirrr/task-management-backend/internal/repository"
)

type TaskService interface {
	CreateTask(userID int, req *model.TaskCreateRequest) (*model.Task, error)
	GetTaskByID(taskID, userID int, isAdmin bool) (*model.Task, error)
	GetUserTasks(userID int, page, limit int, status, search string) (*model.TasksResponse, error)
	GetAllTasks(page, limit int, status, search string) (*model.TasksResponse, error)
	UpdateTask(taskID, userID int, req *model.TaskUpdateRequest, isAdmin bool) (*model.Task, error)
	DeleteTask(taskID, userID int, isAdmin bool) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

// CreateTask membuat task baru
func (s *taskService) CreateTask(userID int, req *model.TaskCreateRequest) (*model.Task, error) {
	task := &model.Task{
		UserID:      userID,
		Title:       req.Title,
		Description: req.Description,
		Status:      model.TaskStatusPending, // Default status
	}

	// Override status if provided
	if req.Status != "" {
		task.Status = req.Status
	}

	err := s.taskRepo.Create(task)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	// Get the created task with timestamps
	createdTask, err := s.taskRepo.GetByID(task.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get created task: %w", err)
	}

	return createdTask, nil
}

// GetTaskByID mengambil task berdasarkan ID dengan authorization check
func (s *taskService) GetTaskByID(taskID, userID int, isAdmin bool) (*model.Task, error) {
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	if task == nil {
		return nil, errors.New(model.ErrTaskNotFound)
	}

	// Check authorization
	if !isAdmin && task.UserID != userID {
		return nil, errors.New(model.ErrForbidden)
	}

	return task, nil
}

// GetUserTasks mengambil tasks milik user dengan pagination dan filter
func (s *taskService) GetUserTasks(userID int, page, limit int, status, search string) (*model.TasksResponse, error) {
	tasks, total, err := s.taskRepo.GetByUserID(userID, page, limit, status, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get user tasks: %w", err)
	}

	return &model.TasksResponse{
		Tasks: tasks,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

// GetAllTasks mengambil semua tasks dengan pagination dan filter (admin only)
func (s *taskService) GetAllTasks(page, limit int, status, search string) (*model.TasksResponse, error) {
	tasks, total, err := s.taskRepo.GetAll(page, limit, status, search)
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}

	return &model.TasksResponse{
		Tasks: tasks,
		Total: total,
		Page:  page,
		Limit: limit,
	}, nil
}

// UpdateTask mengupdate task dengan authorization check
func (s *taskService) UpdateTask(taskID, userID int, req *model.TaskUpdateRequest, isAdmin bool) (*model.Task, error) {
	// Get existing task
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	if task == nil {
		return nil, errors.New(model.ErrTaskNotFound)
	}

	// Check authorization
	if !isAdmin && task.UserID != userID {
		return nil, errors.New(model.ErrForbidden)
	}

	// Update fields if provided
	if req.Title != nil {
		task.Title = *req.Title
	}
	if req.Description != nil {
		task.Description = req.Description
	}
	if req.Status != nil {
		task.Status = *req.Status
	}

	err = s.taskRepo.Update(task)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	// Get updated task
	updatedTask, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return nil, fmt.Errorf("failed to get updated task: %w", err)
	}

	return updatedTask, nil
}

// DeleteTask menghapus task dengan authorization check
func (s *taskService) DeleteTask(taskID, userID int, isAdmin bool) error {
	// Get existing task for authorization check
	task, err := s.taskRepo.GetByID(taskID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}
	if task == nil {
		return errors.New(model.ErrTaskNotFound)
	}

	// Check authorization
	if !isAdmin && task.UserID != userID {
		return errors.New(model.ErrForbidden)
	}

	err = s.taskRepo.Delete(taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}