include .env # read from .env file

OUTPUT_BINARY=blog-api-crud
OUTPUT_DIR=./bin
ENTRY_DIR=./cmd/api

.PHONY: build run clean

build:
	@echo "Building..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(OUTPUT_BINARY) $(ENTRY_DIR)

clean:
	@echo "Cleaning files..."
	@rm -rf $(OUTPUT_DIR)

run: build
	@$(OUTPUT_DIR)/$(OUTPUT_BINARY)


# docker commands
.PHONY: compose/up compose/down connect-shell

compose/up:	
	@echo "Starting containers..."
	docker compose up -d

compose/down:
	@echo "Stopping containers..."
	@docker compose down -v

connect-shell:
	@echo "Connecting to postgres container via shell.."
	@docker exec -it ${BLOG_DB_CONTAINER_NAME} /bin/bash
