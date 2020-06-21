package pgsql

import (
	"context"
	"os"
	"testing"
)

var (
	testDb *Database

	// testing context
	ctx = context.Background()
)

// NOTE: A PostgreSQL instance must be running on db:5432 with
// the below environment configuration for these tests to work
func init() {
	// set environment vars
	os.Setenv("PGHOST", "db")
	os.Setenv("PGUSER", "test")
	os.Setenv("PGPASSWORD", "password")
	os.Setenv("PGDATABASE", "test")
	os.Setenv("PGPORT", "5432")
	os.Setenv("PGSSLMODE", "disable")

	// create and init db
	testDb = NewDatabase()

	// drop and recreate tables so tests have a known starting point
	if _, err := testDb.db.Exec(ctx, "DROP TABLE IF EXISTS counter"); err != nil {
		panic(err)
	}
	testDb.createTables()
}

// setupDbTest deletes all rows in the test tables, so each test
// has a known starting point
func setupDbTest() {
	// delete all rows
	if _, err := testDb.db.Exec(ctx, "DELETE FROM counter"); err != nil {
		panic(err)
	}
}

// getCounterVal gets a value from the counter, or fails the test
func getCounterVal(counterId int, t *testing.T) int {
	query := `SELECT val FROM counter WHERE counter_id = $1`

	var val int
	if err := testDb.db.QueryRow(ctx, query, counterId).Scan(&val); err != nil {
		t.Fatal(err)
	}

	return val
}

func TestIncrementCounter(t *testing.T) {
	setupDbTest()
	counterId := 0

	// expect first insert to zero init counter
	if err := testDb.IncrementCounter(0); err != nil {
		t.Fatal(err)
	}

	got := getCounterVal(counterId, t)
	if got != 0 {
		t.Fatalf("expected counter val: %v, got: %v", 0, got)
	}

	// expect an increment
	if err := testDb.IncrementCounter(0); err != nil {
		t.Fatal(err)
	}

	got = getCounterVal(counterId, t)
	if got != 1 {
		t.Fatalf("expected counter val: %v, got: %v", 1, got)
	}
}

func TestGetCounterVal(t *testing.T) {
	setupDbTest()
	counterId := 8
	counterVal := 1234

	// insert a value at counterId 8
	query := `INSERT INTO counter(counter_id, val) VALUES($1, $2)`
	if _, err := testDb.db.Exec(ctx, query, counterId, counterVal); err != nil {
		t.Fatal(err)
	}

	// get the value
	val, err := testDb.GetCounterVal(counterId)
	if err != nil {
		t.Fatal(err)
	}

	if val != counterVal {
		t.Errorf("expected counter val: %v, got: %v", counterVal, val)
	}
}
