package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yuhengfdada/go-bank/api"
	"github.com/yuhengfdada/go-bank/db"
)

const (
	dbDriver      = "postgres"
	dbCredentials = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbCredentials)
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewSQLStore(conn)
	server := api.NewServer(store)
	server.Start()
}
