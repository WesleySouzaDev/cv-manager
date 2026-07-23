# Variables
GO = go
PROJECT_DIR = $(shell pwd)
GQLGEN = github.com/99designs/gqlgen
DB_URL = mongodb://root:1234@localhost:27017/?authSource=admin
DB_NAME = cv_manager

# Environment
PORT = 8080

# Commands
build:
	@echo "===> (1/2) Building Server..."
	@$(GO) build -o $(PROJECT_DIR)/bin/server $(PROJECT_DIR)/server.go
	@echo "===> (2/2) Project built into: $(PROJECT_DIR)/bin/server"

run: build
	@echo "\n===> (1/1) Running Server..."
	@DB_NAME=$(DB_NAME) DB_URL=$(DB_URL) PORT=$(PORT) $(PROJECT_DIR)/bin/server

generate:
	@echo "===> (1/2) Generating GraphQL Schema..."
	@$(GO) run $(GQLGEN) generate
	@echo "===> (2/2) GraphQL Schema Generated With Success."

clear:
	@echo "===> (1/2) Cleaning Binary Project..."
	@rm -rf $(PROJECT_DIR)/bin
	@echo "===> (2/2) $(PROJECT_DIR)/bin Cleaned Up."


.DEFAULT_GOAL := run

