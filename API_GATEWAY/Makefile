# Load .env file
ifneq (,$(wildcard ./.env))
  include .env
  export
endif

MIGRATIONS_DIR=./db/migrations
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):(DB_PORT)/$(DB_NAME)?sslmode=disable

# Create a new migration file with the given name
# migrate-create name=create_users_table
migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# Apply migrations to the database
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" up

# Rollback the last migration
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_URL)" down
