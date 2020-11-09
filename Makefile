.PHONY: build test format

PROJECT_NAME ?= filecoin-indexer
DOCKER_IMAGE ?= figmentnetworks/${PROJECT_NAME}
DOCKER_TAG   ?= latest

# Build the binary
build:
	go build

# Run tests
test:
	go test -race -cover ./...

# Format the code
format:
	go fmt ./...

# Build a local Docker image
docker:
	docker build -t ${PROJECT_NAME} -f Dockerfile .
