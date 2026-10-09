.PHONY: run build docker-up docker-down sqlc test seed

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go
	go build -o bin/seeder cmd/seeder/main.go

docker-up:
	docker compose up --build

docker-down:
	docker compose down -v

sqlc:
	sqlc generate

test:
	go test -v -race ./tests/...

seed:
	go run cmd/seeder/main.go
