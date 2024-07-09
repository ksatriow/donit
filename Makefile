# Makefile

BINARY_NAME=donit

all: build

build:
	@echo "Building the Go project..."
	@go build -o $(BINARY_NAME)

install: build
	@echo "Installing the binary..."
	@mv $(BINARY_NAME) /usr/local/bin/

clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

.PHONY: all build install clean
