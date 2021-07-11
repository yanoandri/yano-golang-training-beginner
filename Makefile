container-up:
	docker compose up -d

container-down:
	docker compose down

container-rebuild:
	docker compose down && docker compose build && docker compose up -d

migrate-up:
	migrate -database "postgres://test:test@postgres:5432/payment?sslmode=disable" -path ./migrations up

migrate-down:
	migrate -database "postgres://test:test@postgres:5432/payment?sslmode=disable" -path ./migrations down