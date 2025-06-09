package persistence

import (
	"context"
	"sync"

	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/domain/model"
)

// TodoRepository implements the TodoRepositoryPort interface
type TodoRepository struct {
	todos  map[int]*model.Todo
	mu     sync.RWMutex
	nextID int
}

// NewTodoRepository creates a new TodoRepository
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		todos:  make(map[int]*model.Todo),
		nextID: 1,
	}
}

// Create creates a new todo
func (r *TodoRepository) Create(ctx context.Context, todo *model.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	todo.ID = r.nextID
	r.nextID++
	r.todos[todo.ID] = todo
	return nil
}

// GetByID retrieves a todo by ID
func (r *TodoRepository) GetByID(ctx context.Context, id int) (*model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, exists := r.todos[id]
	if !exists {
		return nil, domain.ErrNotFound
	}
	return todo, nil
}

// List retrieves all todos
func (r *TodoRepository) List(ctx context.Context) ([]*model.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todos := make([]*model.Todo, 0, len(r.todos))
	for _, todo := range r.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

// Update updates an existing todo
func (r *TodoRepository) Update(ctx context.Context, todo *model.Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[todo.ID]; !exists {
		return domain.ErrNotFound
	}

	r.todos[todo.ID] = todo
	return nil
}

// Delete deletes a todo
func (r *TodoRepository) Delete(ctx context.Context, id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.todos[id]; !exists {
		return domain.ErrNotFound
	}

	delete(r.todos, id)
	return nil
}
