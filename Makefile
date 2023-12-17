PROJECT_NAME := getclaim
MODULE_NAME := github.com/mamalovesyou/$(PROJECT_NAME)

DEPLOYMENT := $(CURDIR)/deployment
LOCAL_DOCKER_COMPOSE := $(DEPLOYMENT)/local
DOCKER_COMPOSE_CLEAN_FLAGS=--volumes --rmi local --remove-orphans

# Docker compose
DOCKER_COMPOSE_ENV = COMPOSE_DOCKER_CLI_BUILD=1
DOCKER_COMPOSE_CMD = docker-compose -p $(PROJECT_NAME)


devtools:
	@echo Installing tools from tools/tools.go
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install % \
    		&& echo "✅ Tools installed" || (echo "❌ Failed to install tools"; exit 1);

gql:
	go run github.com/99designs/gqlgen gen

########################
###  Docker Compose  ###
########################

dev: ## Start dev environment with docker
	@echo "Starting dev infra..."
	$(DOCKER_COMPOSE_CMD) -f $(LOCAL_DOCKER_COMPOSE)/docker-compose.yml  --env-file .env up

dev.build: ## Start dev environment with docker
	@echo "Starting dev infra..."
	$(DOCKER_COMPOSE_CMD) -f $(LOCAL_DOCKER_COMPOSE)/docker-compose.yml  --env-file .env up --build

dev.clean: ## Clean docker dev evironment
	@echo "Cleaning dev env..."
	$(DOCKER_COMPOSE_CMD) -f $(LOCAL_DOCKER_COMPOSE)/docker-compose.yml down
	$(DOCKER_COMPOSE_CMD) -f $(LOCAL_DOCKER_COMPOSE)/docker-compose.yml rm -f


help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+(\.[a-zA-Z_-]+)*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
