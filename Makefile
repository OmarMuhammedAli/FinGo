postgres:
	docker run --name pgmasterclass -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb: 
	docker exec -it pgmasterclass createdb --username=root --owner=root fingo

createdb-test:
	docker exec -it pgmasterclass createdb --username=root --owner=root fingo-test

dropdb:
	docker exec -it pgmasterclass dropdb fingo

migrate:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo?sslmode=disable" -verbose up

rollback:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo?sslmode=disable" -verbose down 1

rollback_all:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo-test?sslmode=disable" -verbose down

migrate-test:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo-test?sslmode=disable" -verbose up

rollback-test:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo-test?sslmode=disable" -verbose down 1

rollback_all-test:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/fingo?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb createdb-test dropdb migrate rollback rollback_all migrate-test rollback-test rollback_all-test sqlc test