package main

import (
	"database/sql"
	"log"

	"github.com/vinicius-gregorio/simple_bank/api"
	db "github.com/vinicius-gregorio/simple_bank/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
