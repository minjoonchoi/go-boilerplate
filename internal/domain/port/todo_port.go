package port

import (
	"context"
	"go-boilerplate/internal/domain/model"
)

// TodoRepositoryPort defines the interface for todo persistence
type TodoRepositoryPort interface {
	Create(ctx context.Context, todo *model.Todo) error
	GetByID(ctx context.Context, id int) (*model.Todo, error)
	List(ctx context.Context) ([]*model.Todo, error)
	Update(ctx context.Context, todo *model.Todo) error
	Delete(ctx context.Context, id int) error
}

// TodoServicePort defines the interface for todo business logic
type TodoServicePort interface {
	CreateTodo(ctx context.Context, req *model.CreateTodoRequest) (*model.Todo, error)
	GetTodo(ctx context.Context, id int) (*model.Todo, error)
	ListTodos(ctx context.Context) ([]*model.Todo, error)
	UpdateTodo(ctx context.Context, id int, req *model.UpdateTodoRequest) (*model.Todo, error)
	DeleteTodo(ctx context.Context, id int) error
}
