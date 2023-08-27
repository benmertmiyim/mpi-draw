package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:d4ERBLCFmfK10U80Z0o4lGD2s9l9Tlxr@localhost:5432/mpi_draw?sslmode=disable"
)

var TestQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	TestQueries = New(conn)

	os.Exit(m.Run())
}
