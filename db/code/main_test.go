package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

var (
	dbDriver      string = "postgres"
	dbCredentials string = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testStore *Store

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbCredentials)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = New(conn)
	testStore = NewStore(conn)
	m.Run()
}
