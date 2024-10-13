# Variables
DB_URL := postgresql://root:root@localhost:5432/flightDB?sslmode=disable
MIGRATE := $(shell which migrate)
ifeq ($(MIGRATE),)
    ifneq ($(GOPATH),)
        MIGRATE := $(GOPATH)/bin/migrate
    else
        MIGRATE := $(HOME)/go/bin/migrate
    endif
endif

UNAME := $(shell uname)

# Set the open command based on the OS
ifeq ($(UNAME), Darwin)
    OPEN_CMD = open
else ifeq ($(UNAME), Linux)
    OPEN_CMD = xdg-open
else
    OPEN_CMD = start
endif

# File to be opened
FILE_TO_OPEN = coverage.html

# Targets
.PHONY: install-migrate
install-migrate:
	@echo "Installing golang-migrate..."
	@go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: create-migration
create-migration:
	@echo "Creating new migration..."
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir db/migrations -seq $${name}

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	@$(MIGRATE) -database "$(DB_URL)" -path db/migrations up

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back the last migration..."
	@$(MIGRATE) -database "$(DB_URL)" -path db/migrations down 1

.PHONY: migrate-force
migrate-force:
	@echo "Forcing migration version..."
	@read -p "Enter migration version: " version; \
	$(MIGRATE) -database "$(DB_URL)" -path db/migrations force $${version}

.PHONY: migrate-drop
migrate-drop:
	@echo "Dropping all tables..."
	@$(MIGRATE) -database "$(DB_URL)" -path db/migrations drop

.PHONY: migrate-version
migrate-version:
	@echo "Checking current migration version..."
	@$(MIGRATE) -database "$(DB_URL)" -path db/migrations version

.PHONY: build
build:
	@echo "Building the application..."
	@go build -o build/ main.go

.PHONY: run
run: build
	@echo "Running the application..."
	@./build/main
