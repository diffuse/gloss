package gloss

import "os"

// NOTE: A PostgreSQL instance must be running on db:5432 with
// the below environment configuration for these tests to work
func init() {
	// set environment vars
	os.Setenv("POSTGRES_HOST", "db")
	os.Setenv("POSTGRES_USER", "test")
	os.Setenv("POSTGRES_PASS", "password")
	os.Setenv("POSTGRES_DBNAME", "test")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_SSL_MODE", "disable")

	// init db
	InitDb()
}

// setupDbTest drops the table counter if it exists and
// creates new, fresh tables for the test to run with
func setupDbTest() {
	// drop the table if it exists
	if _, err := db.Exec("DROP TABLE IF EXISTS counter"); err != nil {
		panic(err)
	}

	// create fresh tables
	CreateTables()
}