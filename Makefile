
## TEST
test:
	go test ./... -v

test.cover:
	go test ./... --cover

## INFRA
up.database:
	docker compose up postgres -d

## RUN APP
run.api:
	go run cmd/*.go api

## MIGRATION
run.migrations:
	go run cmd/*.go migrations

## SETUP
create.env:
	cp .env.example .env

setup: create.env up.database run.migrations
