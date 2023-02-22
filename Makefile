postgres:
	docker run --name db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -d postgres:14-alpine
createdb:
	docker exec -it db createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it db dropdb simple_bank 
migrateup:
	migrate -path db/migration -database "postgresql://root:1234@db:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:1234@db:5432/simple_bank?sslmode=disable" -verbose down	
migrateup1:
	migrate -path db/migration -database "postgresql://root:1234@db:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown1:
	migrate -path db/migration -database "postgresql://root:1234@db:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server migratedown1 migrateup1