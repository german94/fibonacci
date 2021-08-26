.DEFAULT_GOAL := build
BIN_FILE=fibonacci

build:
	@go build -o "cmd/fibonacci/fibonacci" fibonacci/cmd/fibonacci

run:
	@cd cmd/fibonacci; ./fibonacci

clean:
	@cd cmd/fibonacci; go clean;

test:
	go test ./...