include .env

.PHONY: fmt vet test lint up setup

fmt:
	gofmt -w .

vet:
	go vet ./...

test:
	go test ./...

lint:
	docker run --rm \
		-v $(PWD):/app \
		-w /app \
		--entrypoint golangci-lint \
		golangci/golangci-lint:v2.13.2 \
		run ./...

up:
	docker compose up -d

setup: up
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/schema.sql
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/seed.sql
