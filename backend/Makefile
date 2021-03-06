## Print the help message.
# Parses this Makefile and prints targets that are preceded by "##" comments.
help:
	@echo "" >&2
	@echo "Available targets: " >&2
	@echo "" >&2
	@awk -F : '\
			BEGIN { in_doc = 0; } \
			/^##/ && in_doc == 0 { \
				in_doc = 1; \
				doc_first_line = $$0; \
				sub(/^## */, "", doc_first_line); \
			} \
			$$0 !~ /^#/ && in_doc == 1 { \
				in_doc = 0; \
				if (NF <= 1) { \
					next; \
				} \
				printf "  %-15s %s\n", $$1, doc_first_line; \
			} \
			' <"$(abspath $(lastword $(MAKEFILE_LIST)))" \
		| sort >&2
	@echo "" >&2

## run the server locally as configured by a .env file in the root of backend dir
run:
	go run cmd/main.go

## generates a .env file that can be used to configure local dev
envfile:
	cp .env.example .env

.PHONY: api
SWAGGER_GEN = swagger generate server -q --target ./api --name rollpay --spec ./api/swagger.yml --model-package model -P model.Principal -a operation -s rest mixin --exclude-main
## Generate the API models and server
api:
	$(MAKE) _dexec CMD="$(SWAGGER_GEN)"

## Start development environment (generates code, spins up databases, etc)
dev.start:
	mkdir -p .cache/pkg
	echo "LOCAL_MOUNT=$(shell pwd)" >> .env
	docker compose up -d

## Stop the development environment
dev.stop:
	docker compose down

name = ""
## Creates a new db migration file for the provided service. (Ex: make migration name=cool-new-tables)
migration:
	$(MAKE) _dexec CMD="sql-migrate new -config=db/.migrations/dbconfig.yml -env=rollpay $(name)"

## Run an up migration
migrate.up:
	$(MAKE) _dexec CMD="sql-migrate up -config=db/.migrations/dbconfig.yml -env=rollpay"

## Run a down migration
migrate.down:
	$(MAKE) _dexec CMD="sql-migrate down -config=db/.migrations/dbconfig.yml -env=rollpay"

## Generates database models based upon the existing tables in your local development database
dbmodels:
	rm -rf ./datastore/model/*.go
	$(MAKE) _dexec CMD="sqlboiler -c datastore/model/sqlboiler.toml -o datastore/model -p model --tag db --no-hooks psql --add-soft-deletes"

_dexec:
	docker exec rollpay_dev ${CMD}
