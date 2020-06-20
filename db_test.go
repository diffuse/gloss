package gloss

import (
	"os"
	"testing"
)

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

// getCounterVal gets a value from the counter, or fails the test
func getCounterVal(counterId int, t *testing.T) int {
	query := `SELECT val FROM counter WHERE counter_id = $1`

	var val int
	if err := db.QueryRow(query, counterId).Scan(&val); err != nil {
		t.Fatal(err)
	}

	return val
}

func TestIncrementCounter(t *testing.T) {
	setupDbTest()
	counterId := 0

	// expect first insert to zero init counter
	if err := IncrementCounter(0); err != nil {
		t.Fatal(err)
	}

	got := getCounterVal(counterId, t)
	if got != 0 {
		t.Fatalf("expected counter val: %v, got: %v", 0, got)
	}

	// expect an increment
	if err := IncrementCounter(0); err != nil {
		t.Fatal(err)
	}

	got = getCounterVal(counterId, t)
	if got != 1 {
		t.Fatalf("expected counter val: %v, got: %v", 1, got)
	}
}

func TestGetCounterVal(t *testing.T) {
	// insert a value at counterId 8
	counterId := 8
	counterVal := 1234
	query := `INSERT INTO counter(counter_id, val) VALUES($1, $2)`
	if _, err := db.Exec(query, counterId, counterVal); err != nil {
		t.Fatal(err)
	}

	// get the value
	val, err := GetCounterVal(counterId)
	if err != nil {
		t.Fatal(err)
	}

	if val != counterVal {
		t.Errorf("expected counter val: %v, got: %v", counterVal, val)
	}
}