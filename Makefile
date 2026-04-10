PHONY: createdb dropdb postgres migrate-up migrate-down sqlc test

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12alpine

startPostgres:
	docker start postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simplebank

dropdb:
	docker exec -it postgres12 dropdb simplebank

migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable" --verbose down 

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...
