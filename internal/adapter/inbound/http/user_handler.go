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

// UserHandler handles HTTP requests for users
type UserHandler struct {
	userService port.UserServicePort
}

// NewUserHandler creates a new UserHandler
func NewUserHandler(userService port.UserServicePort) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// mapDomainErrorToHTTP maps domain errors to appropriate HTTP status codes
func (h *UserHandler) mapDomainErrorToHTTP(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound, "User not found"
	case errors.Is(err, domain.ErrInvalidUsername):
		return http.StatusBadRequest, "Invalid username"
	case errors.Is(err, domain.ErrUsernameDuplicate):
		return http.StatusConflict, "Username already exists"
	case errors.Is(err, domain.ErrInvalidEmail):
		return http.StatusBadRequest, "Invalid email format"
	case errors.Is(err, domain.ErrDuplicate):
		return http.StatusConflict, "Resource already exists"
	default:
		return http.StatusInternalServerError, "Internal server error"
	}
}

// CreateUser handles POST /users
// @Summary Create a new user
// @Description Create a new user with username, email, and name
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.CreateUserRequest true "User object"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 409 {object} map[string]string "Conflict - Username already exists"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.CreateUser(c.Request.Context(), &req)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUser handles GET /users/:id
// @Summary Get a user by ID
// @Description Get user information by user ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string "Bad Request - Invalid ID"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), id)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, user)
}

// ListUsers handles GET /users
// @Summary List all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userService.ListUsers(c.Request.Context())
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser handles PUT /users/:id
// @Summary Update a user
// @Description Update user information by user ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.UpdateUserRequest true "User update object"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string "Bad Request"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), id, &req)
	if err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser handles DELETE /users/:id
// @Summary Delete a user
// @Description Delete a user by user ID
// @Tags users
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string "Bad Request - Invalid ID"
// @Failure 404 {object} map[string]string "Not Found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	if err := h.userService.DeleteUser(c.Request.Context(), id); err != nil {
		statusCode, message := h.mapDomainErrorToHTTP(err)
		c.JSON(statusCode, gin.H{"error": message})
		return
	}

	c.Status(http.StatusNoContent)
}
