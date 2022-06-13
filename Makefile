.PHONY: postgres create start stop migrateup migratedown generate

postgres:
	docker pull postgres:latest
	docker run --name postgres14 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres
create:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank
start:
	docker start postgres14
stop:
	docker stop postgres14
psql:
	docker exec -it postgres14 psql -U root -d simple_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

generate:
	sqlc generate

test:
	go test -v ./...
server:
	go run main.go