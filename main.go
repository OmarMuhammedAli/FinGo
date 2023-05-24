package main

import (
	"database/sql"
	"log"

	"github.com/OmarMuhammedAli/FinGo/api"
	db "github.com/OmarMuhammedAli/FinGo/db/sqlc"
	"github.com/OmarMuhammedAli/FinGo/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load configs")
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to db: %v\n", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}
