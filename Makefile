# Go Boilerplate Makefile
# 새로운 프로젝트 생성 및 관리를 위한 명령어들

.PHONY: help create-project run build test docs clean install lint format

# 기본값
PROJECT_NAME ?= go-boilerplate
BINARY_NAME := $(PROJECT_NAME)
CMD_DIR := ./cmd
BUILD_DIR := ./bin
SCRIPTS_DIR := ./scripts

# 색상 코드
BLUE := \033[34m
GREEN := \033[32m
YELLOW := \033[33m
RED := \033[31m
RESET := \033[0m

# 기본 타겟
help: ## 🆘 Show this help message
	@echo ""
	@echo "$(BLUE)  ________      ____        _ __                 __      __$(RESET)"
	@echo "$(BLUE) / ____/ /___  / __ )____  (_) /_  ___  _____  / /_____/ /_$(RESET)"
	@echo "$(BLUE)/ / __/ / __ \\/ __  / __ \\/ / / / _ \\/ ___/ __ \\/ / ___/ __/$(RESET)"
	@echo "$(BLUE)/ /_/ / / /_/ / /_/ / /_/ / / / /  __/ /  / /_/ / / /  / /_$(RESET)"
	@echo "$(BLUE)\\____/_/\\____/_____/\\____/_/_/_/\\___/_/  / .___/_/_/   \\__/$(RESET)"
	@echo "$(BLUE)                                       /_/$(RESET)"
	@echo ""
	@echo "$(GREEN)🚀 Go Boilerplate Commands:$(RESET)"
	@echo ""
	@awk 'BEGIN {FS = ":.*##"; printf ""} /^[a-zA-Z_-]+:.*?##/ { printf "  $(YELLOW)%-18s$(RESET) %s\\n", $$1, $$2 }' $(MAKEFILE_LIST)
	@echo ""

create-project: ## 🎯 Create a new project from this boilerplate
	@echo "$(GREEN)🚀 Creating new project from go-boilerplate template...$(RESET)"
	@chmod +x $(SCRIPTS_DIR)/create-project.sh
	@$(SCRIPTS_DIR)/create-project.sh

run: ## 🏃 Run the development server
	@echo "$(GREEN)🏃 Starting development server...$(RESET)"
	@if [ -f "$(SCRIPTS_DIR)/run.sh" ]; then \
		chmod +x $(SCRIPTS_DIR)/run.sh && $(SCRIPTS_DIR)/run.sh; \
	else \
		go run $(CMD_DIR)/main.go; \
	fi

build: ## 🔨 Build the project
	@echo "$(GREEN)🔨 Building project...$(RESET)"
	@if [ -f "$(SCRIPTS_DIR)/build.sh" ]; then \
		chmod +x $(SCRIPTS_DIR)/build.sh && $(SCRIPTS_DIR)/build.sh; \
	else \
		mkdir -p $(BUILD_DIR) && go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)/main.go; \
	fi

test: ## 🧪 Run all tests
	@echo "$(GREEN)🧪 Running tests...$(RESET)"
	@if [ -f "$(SCRIPTS_DIR)/test.sh" ]; then \
		chmod +x $(SCRIPTS_DIR)/test.sh && $(SCRIPTS_DIR)/test.sh; \
	else \
		go test -v ./...; \
	fi

docs: ## 📚 Generate Swagger documentation
	@echo "$(GREEN)📚 Generating Swagger documentation...$(RESET)"
	@if command -v swag >/dev/null 2>&1; then \
		swag init -g $(CMD_DIR)/main.go; \
		echo "$(GREEN)✅ Swagger docs generated at ./docs/$(RESET)"; \
	elif [ -f "$$HOME/go/bin/swag" ]; then \
		$$HOME/go/bin/swag init -g $(CMD_DIR)/main.go; \
		echo "$(GREEN)✅ Swagger docs generated at ./docs/$(RESET)"; \
	else \
		echo "$(YELLOW)⚠️  Installing Swagger CLI...$(RESET)"; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		$$HOME/go/bin/swag init -g $(CMD_DIR)/main.go; \
		echo "$(GREEN)✅ Swagger docs generated at ./docs/$(RESET)"; \
	fi

clean: ## 🧹 Clean build artifacts
	@echo "$(GREEN)🧹 Cleaning build artifacts...$(RESET)"
	@rm -rf $(BUILD_DIR)
	@rm -rf tmp
	@rm -f *.log
	@go clean
	@echo "$(GREEN)✅ Clean completed$(RESET)"

install: ## 📦 Install project dependencies
	@echo "$(GREEN)📦 Installing dependencies...$(RESET)"
	@go mod download
	@go mod tidy
	@echo "$(GREEN)✅ Dependencies installed$(RESET)"

lint: ## 🔍 Run linter (requires golangci-lint)
	@echo "$(GREEN)🔍 Running linter...$(RESET)"
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "$(YELLOW)⚠️  golangci-lint not found. Install it with:$(RESET)"; \
		echo "$(BLUE)curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.54.2$(RESET)"; \
	fi

format: ## 🎨 Format Go code
	@echo "$(GREEN)🎨 Formatting code...$(RESET)"
	@go fmt ./...
	@echo "$(GREEN)✅ Code formatted$(RESET)"

dev: install docs run ## 🔧 Setup development environment and run

quick-start: ## ⚡ Quick start guide
	@echo ""
	@echo "$(GREEN)⚡ Quick Start Guide:$(RESET)"
	@echo ""
	@echo "$(YELLOW)1. Create a new project:$(RESET)"
	@echo "   $(BLUE)make create-project$(RESET)"
	@echo ""
	@echo "$(YELLOW)2. Run development server:$(RESET)"
	@echo "   $(BLUE)make run$(RESET)"
	@echo ""
	@echo "$(YELLOW)3. Build project:$(RESET)"
	@echo "   $(BLUE)make build$(RESET)"
	@echo ""
	@echo "$(YELLOW)4. View API docs:$(RESET)"
	@echo "   $(BLUE)Open http://localhost:8080/swagger/index.html$(RESET)"
	@echo ""

# 디버그용 - 현재 설정 출력
debug: ## 🐛 Show current configuration
	@echo "$(GREEN)🐛 Current Configuration:$(RESET)"
	@echo "$(BLUE)Project Name:$(RESET) $(PROJECT_NAME)"
	@echo "$(BLUE)Binary Name:$(RESET) $(BINARY_NAME)"
	@echo "$(BLUE)Build Dir:$(RESET) $(BUILD_DIR)"
	@echo "$(BLUE)CMD Dir:$(RESET) $(CMD_DIR)"
	@echo "$(BLUE)Scripts Dir:$(RESET) $(SCRIPTS_DIR)"
	@echo "$(BLUE)Go Version:$(RESET) $(shell go version)"
	@echo "$(BLUE)Working Dir:$(RESET) $(shell pwd)" 