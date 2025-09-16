package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// ⚠️ Update this DSN to match your docker-compose `.env.postgres`
	// Example: postgres://user:password@localhost:5432/dbname?sslmode=disable
	dbSource := "postgres://chinsiang:chinsiang@localhost:5432/go-project?sslmode=disable"

	var err error
	testDB, err = sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Initialize the sqlc Queries object
	testQueries = New(testDB)

	// Run migrations (assuming `migrate` CLI is available)
	// If not, you can use golang-migrate library here
	// Example:
	// err = runMigrations(dbSource, "internal/database/migrations")
	// if err != nil {
	//     log.Fatal("cannot run migrations:", err)
	// }

	// Run tests
	code := m.Run()

	// Clean up if needed
	_ = testDB.Close()

	os.Exit(code)
}
