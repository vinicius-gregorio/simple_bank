postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d postgres

createdb: 
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb: 
	docker exec -it postgres dropdb --username=postgres simple_bank

migrationsup:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrationsdown:
	migrate -path db/migrations -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down


sqlcgenerate:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...	

server: 
	go run main.go

watch:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go

mockgenerate:
	mockgen -package mockdb -destination db/mock/store.go github.com/vinicius-gregorio/simple_bank/db/sqlc Store 

.PHONY: postgres createdb dropdb migrationsup migrationsdown sqlcgenerate test server mockgenerate