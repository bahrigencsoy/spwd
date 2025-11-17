# Makefile for shortpath

# Binary name
BINARY_NAME=shortpath

# Build directory
BUILD_DIR=bin

# Go parameters
GOCMD=go
GOBUILD=\$(GOCMD) build
GOCLEAN=\$(GOCMD) clean
GOTEST=\$(GOCMD) test
GOGET=\$(GOCMD) get
GOMOD=\$(GOCMD) mod

# Build flags for optimization
# -ldflags="-s -w" strips debug information and symbol table
# -trimpath removes file system paths from the binary
BUILD_FLAGS=-ldflags="-s -w" -trimpath

# Optimization flags
# -gcflags for compiler optimizations
# -N disables optimizations (don't use for production)
# -l disables inlining (don't use for production)
# Default: enable all optimizations
OPTIMIZE_FLAGS=-gcflags="all=-l=4"

.PHONY: all build clean test coverage bench install uninstall help

all: test build

## build: Build the binary with optimizations
build:
	@echo "Building \$(BINARY_NAME)..."
	@mkdir -p \$(BUILD_DIR)
	@\$(GOBUILD) \$(BUILD_FLAGS) \$(OPTIMIZE_FLAGS) -o \$(BUILD_DIR)/\$(BINARY_NAME) .
	@echo "Build complete: \$(BUILD_DIR)/\$(BINARY_NAME)"

## build-release: Build optimized release binary
build-release:
	@echo "Building release version of \$(BINARY_NAME)..."
	@mkdir -p \$(BUILD_DIR)
	@CGO_ENABLED=0 \$(GOBUILD) \$(BUILD_FLAGS) -o \$(BUILD_DIR)/\$(BINARY_NAME) .
	@echo "Release build complete: \$(BUILD_DIR)/\$(BINARY_NAME)"

## test: Run tests
test:
	@echo "Running tests..."
	@\$(GOTEST) -v ./...

## coverage: Run tests with coverage
coverage:
	@echo "Running tests with coverage..."
	@\$(GOTEST) -v -coverprofile=coverage.out ./...
	@\$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## bench: Run benchmarks
bench:
	@echo "Running benchmarks..."
	@\$(GOTEST) -bench=. -benchmem ./...

## clean: Clean build files
clean:
	@echo "Cleaning..."
	@\$(GOCLEAN)
	@rm -rf \$(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

## install: Install the binary to /usr/local/bin
install: build
	@echo "Installing \$(BINARY_NAME) to /usr/local/bin..."
	@sudo cp \$(BUILD_DIR)/\$(BINARY_NAME) /usr/local/bin/
	@sudo chmod +x /usr/local/bin/\$(BINARY_NAME)
	@echo "Installation complete"

## uninstall: Remove the binary from /usr/local/bin
uninstall:
	@echo "Uninstalling \$(BINARY_NAME)..."
	@sudo rm -f /usr/local/bin/\$(BINARY_NAME)
	@echo "Uninstall complete"

## deps: Download dependencies
deps:
	@echo "Downloading dependencies..."
	@\$(GOMOD) download
	@\$(GOMOD) tidy

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' \${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'