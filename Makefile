include .env

.PHONY: fmt vet test lint run setup

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

run:
	docker run --name orch-postgres \
    	-e POSTGRES_USER=orch \
    	-e POSTGRES_PASSWORD=password \
    	-e POSTGRES_DB=orch \
    	-p 5432:5432 \
    	-d postgres

setup:
	docker compose up -d
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/schema.sql
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/seed.sql
