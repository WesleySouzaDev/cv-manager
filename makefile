.PHONY: default run build test docs clean
# variables
APP_NAME=cv-manager

default: run

run:
	go run server.go
schema:
	go tool gqlgen generate
build:
	go build -o $(APP_NAME)
test:
	go test -v ./...
docs:
	go run github.com/99designs/gqlgen docs
clean:
	rm -rf $(APP_NAME)
