#!/bin/bash

# 빌드 스크립트
echo "Building go-boilerplate..."

# swag CLI 설치 확인 및 설치
SWAG_PATH="$HOME/go/bin/swag"
if [ ! -f "$SWAG_PATH" ]; then
    echo "Installing swag CLI..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Swagger 문서 생성
echo "Generating Swagger documentation..."
"$SWAG_PATH" init -g cmd/main.go

# Go 빌드
go build -o bin/go-boilerplate ./cmd/main.go

echo "Build completed!" 