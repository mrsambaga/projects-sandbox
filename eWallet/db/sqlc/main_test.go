package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

var (
	testQueries *Queries
	dbSource string
	testDB *sql.DB
)

func TestMain(m *testing.M) {
	loadEnv()
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func loadEnv() {	
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbSource = os.Getenv("DB_URL")
}