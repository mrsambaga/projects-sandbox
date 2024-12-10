package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mrsambaga/projects-sandbox/eWallet/api"
	db "github.com/mrsambaga/projects-sandbox/eWallet/db/sqlc"
)

const (
	dbDriver = "postgres"
	serverAddress = "0.0.0.0:8080"
)

var (
	dbSource string
)

func main() {
	loadEnv()
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func loadEnv() {	
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSource = os.Getenv("DB_URL")
}