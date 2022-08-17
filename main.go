package main

import (
	"database/sql"
	"log"

	"github.com/vinicius-gregorio/simple_bank/api"
	"github.com/vinicius-gregorio/simple_bank/util"

	db "github.com/vinicius-gregorio/simple_bank/db/sqlc"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load configurations", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error starting the server", err)
	}
}
