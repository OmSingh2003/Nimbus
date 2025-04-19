package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	// Open a connection pool to the database.
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to the database:", err)
	}
	defer conn.Close() // Ensure the connection pool is closed when TestMain exits.

	testQueries = New(conn)

	// Run all tests in the package. os.Exit passes the test result code back.
	os.Exit(m.Run())
}
