package http

import (
	"errors"
	"net/http"
	"strconv"

	"go-boilerplate/internal/domain"
	"go-boilerplate/internal/domain/model"
	"go-boilerplate/internal/domain/port"

	"github.com/gin-gonic/gin"
)

// TodoHandler handles HTTP requests for todos
type TodoHandler struct {
	todoService port.TodoServicePort
}

// NewTodoHandler creates a new TodoHandler
func NewTodoHandler(todoService port.TodoServicePort) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

// mapDomainErrorToHTTP maps domain errors to appropriate HTTP status codes
func (h *TodoHandler) mapDomainErrorToHTTP(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, "Resource not found"
	case errors.Is(err, domain.ErrInvalidTodoTitle):
		return http.StatusBadRequest, "Invalid todo title"
	case errors.Is(err, domain.ErrTodoAlreadyCompleted):
		return http.StatusConflict, "Todo is already completed"
	case errors.Is(err, domain.ErrDuplicate):
		return http.StatusConflict, "Resource already exists"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}

// CreateTodo handles POST /todos
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body model.CreateTodoRequest true "Todo object"
// @Success 201 {object} model.Todo
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req model.CreateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.todoService.CreateTodo(c.Request.Context(), &req)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodo handles GET /todos/:id
// @Summary Get a todo
// @Description Get a todo by ID
// @Tags todos
// @Produce json
// @Param id path int true "Todo ID"
// @Success 200 {object} model.Todo
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	todo, err := h.todoService.GetTodo(c.Request.Context(), id)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// ListTodos handles GET /todos
// @Summary List todos
// @Description Get all todos
// @Tags todos
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [get]
func (h *TodoHandler) ListTodos(c *gin.Context) {
	todos, err := h.todoService.ListTodos(c.Request.Context())
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, todos)
}

// UpdateTodo handles PUT /todos/:id
// @Summary Update a todo
// @Description Update a todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path int true "Todo ID"
// @Param todo body model.UpdateTodoRequest true "Todo object"
// @Success 200 {object} model.Todo
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req model.UpdateTodoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.todoService.UpdateTodo(c.Request.Context(), id, &req)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo handles DELETE /todos/:id
// @Summary Delete a todo
// @Description Delete a todo by ID
// @Tags todos
// @Param id path int true "Todo ID"
// @Success 204 "No Content"
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.todoService.DeleteTodo(c.Request.Context(), id); err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.Status(http.StatusNoContent)
}
