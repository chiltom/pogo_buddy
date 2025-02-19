BUILD_DIR=./tmp
SRC_DIR=./cmd
TEST_DIR=./...

.PHONY: build run test fmt lint clean tidy all

build:
	@echo "Building the application..."
	go build -o ${BUILD_DIR}/main $(SRC_DIR)/main.go

run:
	@echo "Running the application..."
	air -c .air.toml

test:
	@echo "Running tests..."
	go test -v -cover $(TEST_DIR)

fmt:
	@echo "Formatting the code..."
	go fmt $(TEST_DIR)

lint:
	@echo "Linting the code..."
	go vet $(TEST_DIR)

clean:
	@echo "Cleaning up..."
	rm -f ${BUILD_DIR}/main

tidy:
	@echo "Tidying up..."
	go mod tidy

all: fmt build test