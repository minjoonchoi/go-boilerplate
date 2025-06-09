package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "go-boilerplate/docs"
	"go-boilerplate/internal/adapter/inbound/http"
	"go-boilerplate/internal/adapter/outbound/persistence"
	"go-boilerplate/internal/config"
	"go-boilerplate/internal/domain/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Boilerplate API
// @version 1.0
// @description This is a sample go boilerplate API server.
// @host localhost:8080
// @BasePath /
func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize repositories
	todoRepo := persistence.NewTodoRepository()
	userRepo := persistence.NewUserRepository()

	// Initialize services
	todoService := service.NewTodoService(todoRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	todoHandler := http.NewTodoHandler(todoService)
	userHandler := http.NewUserHandler(userService)

	// Initialize router
	r := initializeRouter(todoHandler, userHandler)

	// Start server
	go func() {
		log.Printf("Server starting on %s", cfg.Server.Address)
		if err := r.Run(cfg.Server.Address); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	log.Println("Server exited properly")
}

// initializeRouter sets up all routes and middleware
func initializeRouter(todoHandler *http.TodoHandler, userHandler *http.UserHandler) *gin.Engine {
	r := gin.Default()

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Todo routes
	todos := r.Group("/todos")
	{
		todos.POST("", todoHandler.CreateTodo)
		todos.GET("", todoHandler.ListTodos)
		todos.GET("/:id", todoHandler.GetTodo)
		todos.PUT("/:id", todoHandler.UpdateTodo)
		todos.DELETE("/:id", todoHandler.DeleteTodo)
	}

	// User routes
	users := r.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.ListUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	return r
}
