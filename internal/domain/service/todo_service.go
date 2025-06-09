package service

import (
	"context"
	"strings"

	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/domain/model"
	"go-boilerplate/internal/domain/port"
)

// TodoService implements the TodoServicePort interface
type TodoService struct {
	repo port.TodoRepositoryPort
}

// NewTodoService creates a new TodoService
func NewTodoService(repo port.TodoRepositoryPort) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

// CreateTodo creates a new todo
func (s *TodoService) CreateTodo(ctx context.Context, req *model.CreateTodoRequest) (*model.Todo, error) {
	// Business logic validation
	if strings.TrimSpace(req.Title) == "" {
		return nil, domain.ErrInvalidTodoTitle
	}

	todo := &model.Todo{
		Title:       req.Title,
		Description: req.Description,
		Completed:   false,
	}

	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// GetTodo retrieves a todo by ID
func (s *TodoService) GetTodo(ctx context.Context, id int) (*model.Todo, error) {
	return s.repo.GetByID(ctx, id)
}

// ListTodos retrieves all todos
func (s *TodoService) ListTodos(ctx context.Context) ([]*model.Todo, error) {
	return s.repo.List(ctx)
}

// UpdateTodo updates an existing todo
func (s *TodoService) UpdateTodo(ctx context.Context, id int, req *model.UpdateTodoRequest) (*model.Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Business logic validation
	if req.Title != "" {
		if strings.TrimSpace(req.Title) == "" {
			return nil, domain.ErrInvalidTodoTitle
		}
		todo.Title = req.Title
	}

	if req.Description != "" {
		todo.Description = req.Description
	}

	// Business logic: prevent completing already completed todos
	if req.Completed && todo.Completed {
		return nil, domain.ErrTodoAlreadyCompleted
	}
	todo.Completed = req.Completed

	if err := s.repo.Update(ctx, todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// DeleteTodo deletes a todo
func (s *TodoService) DeleteTodo(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
