run: gen-proto
	docker compose --env-file ./config/.env up -d
	go run cmd/debez/main.go -env ./config/.env

migrate-up:
	go run cmd/migrate/main.go -env ./config/.env -command up -path ./migrations

migrate-down:
	go run cmd/migrate/main.go -env ./config/.env -command down -path ./migrations

migrate-version:
	go run cmd/migrate/main.go -env ./config/.env -command version -path ./migrations

gen-proto:
	buf generate