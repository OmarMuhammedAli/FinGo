package main

import (
	"database/sql"
	"log"

	"github.com/OmarMuhammedAli/FinGo/api"
	db "github.com/OmarMuhammedAli/FinGo/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/fingo?sslmode=disable"
	address  = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("Cannot connect to db: %v\n", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(address)
	if err != nil {
		log.Fatalf("Cannot start the server: %v\n", err)
	}
}