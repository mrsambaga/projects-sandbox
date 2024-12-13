package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/mrsambaga/projects-sandbox/eWallet/util"
)

var (
	testQueries *Queries
	testDB *sql.DB
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	config, err := util.LoadConfig("../..");
	if err != nil {
		log.Fatal("Error getting env: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}