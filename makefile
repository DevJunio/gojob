.PHONY: default run build test docs clean

APP_NAME = gojob

default: run

run: build
	@./$(APP_NAME)

build: clean
	@go build -o $(APP_NAME) main.go

test:
	@go test ./...

docs: clean-all
	@swag init

clean:
	@rm -f $(APP_NAME)

clean-all: clean
	@rm -rf docs
