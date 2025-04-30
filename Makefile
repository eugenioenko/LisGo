# Variables
APP_NAME=autentico

# Default target
all: build

# Build the Go binary
build:
		GOOS=js GOARCH=wasm go build -o live/lisgo.wasm wasm/lisgo.go

# Run the application
run:
	go run main.go

# Format code using gofmt
fmt:
	gofmt -w .

# Run tests
.PHONY: test
test:
	go test ./...

# Run a linter (requires `golangci-lint`)
lint:
	golangci-lint run ./...
