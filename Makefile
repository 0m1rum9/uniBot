include .env
export
POSTGRES_URL=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable



run:
	go run ./main.go

migrate-up:
	goose postgres "$(POSTGRES_URL)" up -dir ./internal/db/migrations/

docker-up:
	docker-compose up -d --build

