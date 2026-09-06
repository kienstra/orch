include .env

.PHONY: run
run:
	docker run --name orch-postgres \
      -e POSTGRES_USER=orch \
      -e POSTGRES_PASSWORD=password \
      -e POSTGRES_DB=orch \
      -p 5432:5432 \
      -d postgres

.PHONY: setup
setup:
	docker compose up -d
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/schema.sql
	docker compose exec -T postgres psql -U $(POSTGRES_USER) -d $(POSTGRES_DB) < db/seed.sql
