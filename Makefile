postgres:
	docker run --name postgres  -p 5432:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d postgres

createdb: 
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank

dropdb: 
	docker exec -it postgres dropdb --username=postgres simple_bank

.PHONY: postgres createdb dropdb