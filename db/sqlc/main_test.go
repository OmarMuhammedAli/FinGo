package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/OmarMuhammedAli/FinGo/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Failed to read configs")
	}
	os.Setenv("ENV", "test")
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to db: %v", err)
	}

	testQueries = New(testDb)

	os.Exit(m.Run())
}
