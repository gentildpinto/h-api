DATABASE_HOST     ?= localhost
DATABASE_PORT     ?= $(shell grep "DATABASE_PORT" .env | cut -d '=' -f2)
DATABASE_NAME 	  ?= $(shell grep "DATABASE_NAME" .env | cut -d '=' -f2)
DATABASE_USERNAME ?= $(shell grep "DATABASE_USERNAME" .env | cut -d '=' -f2)
DATABASE_PASSWORD ?= $(shell grep "DATABASE_PASSWORD" .env | cut -d '=' -f2)
DATABSE_DSN       ?= ${DATABASE_USERNAME}:${DATABASE_PASSWORD}@${DATABASE_HOST}:${DATABASE_PORT}/${DATABASE_NAME}

# Version of migrations - this is optionally used on goto command
V?=

# Number of migrations - this is optionally used on up and down commands
N?=

.PHONY: migrate_setup migrate_up migrate_down migrate_goto migrate_drop_db

migrate_setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing golang-migrate..."; go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest; fi

migrate_up: migrate_setup
	@ migrate -database 'postgres://${DATABSE_DSN}?sslmode=disable' -path $$(pwd)/migrations up $(N)

migrate_down: migrate_setup
	@ migrate -database 'postgres://${DATABSE_DSN}?sslmode=disable' -path $$(pwd)/migrations down $(N)

migrate_goto: migrate_setup
	@ migrate -database 'postgres://${DATABSE_DSN}?sslmode=disable' -path $$(pwd)/migrations goto $(V)

migrate_drop_db: migrate_setup
	@ migrate -database 'postgres://${DATABSE_DSN}?sslmode=disable' -path $$(pwd)/migrations drop

docker_dev:
	docker compose -f docker-compose.dev.yml up --build
