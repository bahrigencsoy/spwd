# Makefile for the shortpath project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
BINARY_NAME=shortpath
SOURCES=main.go shortpath.go

# Build flags for optimization
LDFLAGS=-ldflags="-s -w"

.PHONY: all build test clean install

all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	$(GOBUILD) -o $(BINARY_NAME) $(LDFLAGS) $(SOURCES)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Clean up build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install the binary
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@mv $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

# Run the application (for development)
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

