.PHONY: all build clean test lint run

# 构建目标
BINARY_NAME=myproject
BUILD_DIR=./bin

all: clean build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/myapp

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Linting..."
	@golangci-lint run

run:
	@echo "Running application..."
	@go run ./cmd/myapp/main.go
