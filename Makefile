.PHONY: build test format

# Build the binary
build:
	go build

# Run tests
test:
	go test -race -cover ./...

# Format the code
format:
	go fmt ./...
