package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mrsambaga/projects-sandbox/eWallet/api"
	db "github.com/mrsambaga/projects-sandbox/eWallet/db/sqlc"
	"github.com/mrsambaga/projects-sandbox/eWallet/util"
)

func main() {
	config, err := util.LoadConfig(".");
	if err != nil {
		log.Fatal("Failed getting env file: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}