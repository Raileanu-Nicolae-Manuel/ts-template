include .env

DB_URL="mysql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?parseTime=true&multiStatements=true"
MIGRATION_FOLDR="file://$(MIGRATIONS_PATH)"

# run need to have the last verion of sqlc generated the migrations will run automatically at start
.PHONY: run
run: sqlc
	go run main.go

.PHONY: sqlc
sqlc:
	sqlc generate

.PHONY: migrate-create
migrate-create:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "Usage: make migrate-create <migration-name>"; \
		exit 1; \
	fi
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) up

.PHONY: migrate-down
migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) down

.PHONY: migrate-version
migrate-version:
	migrate -database $(DB_URL) -path $(MIGRATIONS_PATH) version

.PHONY: proto
proto:
	protoc --proto_path=../proto ../proto/*.proto --go_out=. --go-grpc_out=.
