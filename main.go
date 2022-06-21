package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/yuhengfdada/go-bank/api"
	"github.com/yuhengfdada/go-bank/db"
	"github.com/yuhengfdada/go-bank/util"
)

func main() {
	config, err := util.ReadConfigFromPath(".")

	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBCredentials)
	if err != nil {
		log.Fatal(err)
		return
	}
	store := db.NewSQLStore(conn)
	server := api.NewServer(store, config)
	server.Start(config.ServerAddr)
}
