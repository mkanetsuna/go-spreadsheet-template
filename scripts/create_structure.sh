#!/bin/bash

# ベースディレクトリとファイルの作成
mkdir -p cmd/setup_scheduler
mkdir -p internal/handlers
mkdir -p internal/scheduler
mkdir -p pkg/health
mkdir -p .github/workflows

# ファイルの作成
touch cmd/main.go
touch cmd/setup_scheduler/main.go
touch internal/handlers/handlers.go
touch internal/scheduler/scheduler.go
touch pkg/health/health.go
touch .github/workflows/deploy.yml
touch Dockerfile
touch go.mod

echo "ディレクトリとファイルが作成されました。"