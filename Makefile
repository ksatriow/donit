# Makefile

BINARY_NAME=donit
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

LDFLAGS := -X 'donit/pkg/version.Version=$(VERSION)' \
           -X 'donit/pkg/version.BuildTime=$(BUILD_TIME)' \
           -X 'donit/pkg/version.GitCommit=$(GIT_COMMIT)' \
           -s -w

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@echo "Version: $(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@go build -ldflags "$(LDFLAGS)" -o $(BINARY_NAME) .

install: build
	@echo "Installing $(BINARY_NAME)..."
	@sudo mv $(BINARY_NAME) /usr/local/bin/
	@echo "✓ $(BINARY_NAME) installed successfully!"

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "✓ Cleanup complete"

test:
	@echo "Running tests..."
	@go test -v ./...

.PHONY: all build install clean test
