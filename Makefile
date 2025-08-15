APP_NAME := gotcha
BUILD_DIR := bin

.PHONY: all
all: build

.PHONY: build-static
build-static:
	env CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) -a -ldflags '-extldflags "-static"' .

.PHONY: build
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) .

.PHONY: run
run:
	go run .

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	go vet ./...

