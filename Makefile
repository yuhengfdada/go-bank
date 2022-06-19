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
migrateup1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/yuhengfdada/go-bank/db Store
generate:
	sqlc generate
	make mock


test:
	go test -v ./...
server:
	go run main.go

swagger:
	swagger generate spec -o ./swagger.yaml