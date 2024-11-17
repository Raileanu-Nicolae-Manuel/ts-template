include .env

DB_URL="mysql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?parseTime=true&multiStatements=true"
MIGRATION_FOLDR="file://$(MIGRATIONS_PATH)?format=golang-migrate"

.PHONY: all
all: atlas-diff 

.PHONY: combine-schemas
combine-schemas:
	@echo "Combining schemas..."
	@cat $(SCHEMA_FOLDER)/*.sql > $(SCHEMA_PATH)
# Atlas migration commands

.PHONY: atlas-diff 
atlas-diff: create-db combine-schemas
	@read -p "Enter migration name: " name; \
	atlas migrate diff $${name} \
	--dir $(MIGRATION_FOLDR) \
	--to "file://$(SCHEMA_PATH)" \
	--dev-url $(DB_URL)

.PHONY: create-db
create-db:
	@echo "Checking if database $(DB_NAME) exists..."; \
	if ! mysql -u $(DB_USER) -p$(DB_PASSWORD) -h $(DB_HOST) -P $(DB_PORT) -e "SHOW DATABASES LIKE '$(DB_NAME)';" | grep -q $(DB_NAME); then \
		make atlas-hash; \
		echo "Database $(DB_NAME) not found. Creating..."; \
		mysql -u $(DB_USER) -p$(DB_PASSWORD) -h $(DB_HOST) -P $(DB_PORT) -e "CREATE DATABASE $(DB_NAME);"; \
	fi;

.PHONY: atlas-hash
atlas-hash:
	atlas migrate hash --dir $(MIGRATION_FOLDR)

.PHONY: atlas-apply
atlas-apply:
	atlas migrate apply \
	--dir $(MIGRATION_FOLDR) \
	--url $(DB_URL) \
	--allow-dirty

# status
.PHONY: atlas-status
atlas-status:
	atlas migrate status \
	--dir $(MIGRATION_FOLDR) \
	--url $(DB_URL)



