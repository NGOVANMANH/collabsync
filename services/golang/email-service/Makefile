# Name of your Go binary
BINARY_NAME=email-service

# Go source directory
MAIN_PATH=cmd/main.go

# Default target
.PHONY: run build clean tidy

# Run the service
run:
	go run $(MAIN_PATH)

# Build the binary
build:
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Remove built binary
clean:
	rm -rf bin/

# Tidy up dependencies
tidy:
	go mod tidy
