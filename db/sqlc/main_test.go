package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbSource   = "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable"
	driverName = "postgres"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()
	if err = testDB.Ping(); err != nil {
		log.Fatal("cannot ping to db", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
