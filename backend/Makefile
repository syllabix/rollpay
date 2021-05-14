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

## Start development environment (generates code, spins up databases, etc)
dev.start:
	echo "LOCAL_MOUNT=$(shell pwd)" > .env
	docker compose up -d

## Stop the development environment
dev.stop:
	docker compose down

name = ""
## Creates a new db migration file for the provided service. (Ex: make migration name=cool-new-tables)
migration:
	$(MAKE) _dexec CMD="sql-migrate new -config=.migrations/dbconfig.yml -env=rollpay $(name)"