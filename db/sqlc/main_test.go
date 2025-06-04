package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testStore *Store


func TestMain(m *testing.M) {
	// Open a connection pool to the database
	config, err := util.LoadConfig("../..")
	if err != nill {
		log.Fatal("Cannot load config file",err)
	}
	var err error 
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
