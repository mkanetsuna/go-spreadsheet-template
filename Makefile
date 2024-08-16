.PHONY: dev build run test install-tools

dev:
	@if ! command -v air > /dev/null; then \
		echo "Installing air..."; \
		go install github.com/air-verse/air@latest; \
	fi
	@air

build:
	@go build -o bin/server ./cmd

run: build
	@./bin/server

test:
	@go test ./...

install-tools:
	@go install github.com/air-verse/air@latest