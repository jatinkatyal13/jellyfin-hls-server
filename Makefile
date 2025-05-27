PROJECT_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: default
default: help

######################### SECTION START/STOP #########################
start:
	@echo "\n--------------------> BEGIN : $(section) <--------------------\n"
stop:
	@echo "\n--------------------> DONE : $(section) <--------------------\n"


######################### BINARY #########################
APP_BINARY_DIR=dist
APP_BINARY_PATH=$(APP_BINARY_DIR)/server

## Build
build/clean: ## clean target binary
	@rm -rf $(APP_BINARY_PATH)

build: ## build the binary for local run
	@make start section="BINARY BUILD"
	@rm -f $(APP_BINARY_PATH)
	@go build -o $(APP_BINARY_PATH) ./main.go
	@make stop section="BINARY BUILD"

build/run: build ## build and run the binary using env vars from builder/local_env_vars.env
	@make start section="BINARY RUN"
	@env $(APP_BINARY_PATH)
	@make stop section="BINARY RUN"


######################### DEVELOPMENT #########################
DOCKER_IMAGE_NAME=jellyfin-hls-server
DOCKER_FILE_PATH=Dockerfile
DOCKER_COMPOSE_FILE_PATH=dev/docker-compose.yaml

## Development
dev/start: ## start the development environment
	@make start section="DEV"
	@docker compose -f $(DOCKER_COMPOSE_FILE_PATH) up --build
	@make stop section="DEV"

dev/ps: ## list the containers in the development environment
	@docker compose -f $(DOCKER_COMPOSE_FILE_PATH) ps

dev/stop: ## stop the development environment
	@make start section="DEV STOP"
	@docker compose -f $(DOCKER_COMPOSE_FILE_PATH) down
	@make stop section="DEV STOP"


######################### MIGRATION #########################
MIGRATION_DIR=migrations

## DBMate
dbmate/status: ## show the status of the database migrations
	@make start section="DBMATE"
	@dbmate -d $(MIGRATION_DIR) status
	@make stop section="DBMATE"

dbmate/up: ## apply the database migrations
	@make start section="DBMATE"
	@dbmate -d $(MIGRATION_DIR) up
	@make stop section="DBMATE"

dbmate/down: ## rollback the last database migration
	@make start section="DBMATE"
	@dbmate -d $(MIGRATION_DIR) down
	@make stop section="DBMATE"

dbmate/new: ## create a new database migration
	@make start section="DBMATE NEW"
	@read -p "Enter migration name: " name; \
	if [ -z "$$name" ]; then \
		echo "Migration name cannot be empty"; \
		exit 1; \
	fi; \
	dbmate -d $(MIGRATION_DIR) new "$$name"
	@make stop section="DBMATE NEW"


############################# HELP #####################################
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Help:
help: ## show this help
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} { \
		if (/^[a-zA-Z\/_-]+:.*?##.*$$/) {printf "    ${YELLOW}%-30s${GREEN}%s${RESET}\n", $$1, $$2} \
		else if (/^## .*$$/) {printf "  ${CYAN}%s${RESET}\n", substr($$1,4)} \
		}' $(MAKEFILE_LIST)
