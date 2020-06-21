package gloss

import (
	"database/sql"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var db *sql.DB

// InitDb connects to a PostgreSQL instance and creates the
// tables this service relies on if they don't already exist
func InitDb() {
	// connect to database using configuration created from environment variables
	//
	// from https://godoc.org/github.com/lib/pq
	// "Most environment variables as specified at http://www.postgresql.org/docs/current/static/libpq-envars.html
	// supported by libpq are also supported by pq."
	connector, err := pq.NewConnector("")
	if err != nil {
		panic(err)
	}

	db = sql.OpenDB(connector)

	// create tables
	CreateTables()
}

// CloseDb closes connections to the database
func CloseDb() error {
	return db.Close()
}

// CreateTables creates the tables that this service relies on
func CreateTables() {
	query := `
	CREATE TABLE IF NOT EXISTS counter
	(
		counter_id INTEGER PRIMARY KEY,
		val INTEGER NOT NULL
	)`

	if _, err := db.Exec(query); err != nil {
		panic(err)
	}
}

// IncrementCounter increments the value in the counter table
func IncrementCounter(counterId int) error {
	query := `
	INSERT INTO counter(counter_id, val)
	VALUES($1, 0)
	ON CONFLICT(counter_id)
	DO UPDATE
	SET val = counter.val + 1`

	_, err := db.Exec(query, counterId)
	return err
}

// GetCounterVal gets the value of a counter with ID counterId
func GetCounterVal(counterId int) (int, error) {
	query := `SELECT val FROM counter WHERE counter_id = $1`

	var val int
	return val, db.QueryRow(query, counterId).Scan(&val)
}
