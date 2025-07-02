package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/OmSingh2003/nimbus/util"
	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testStore   Store
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config file:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	defer testDB.Close() // Ensure the connection pool is closed when TestMain exits.

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	// Run all tests in the package. os.Exit passes the test result code back.
	os.Exit(m.Run())
}
