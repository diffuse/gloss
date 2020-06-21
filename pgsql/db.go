package pgsql

import (
	"database/sql"
	_ "github.com/jackc/pgx"
	"github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

// NewDatabase creates and initializes a new Database
func NewDatabase() *Database {
	db := &Database{}
	db.Init()

	return db
}

// Init connects to a PostgreSQL instance and creates the
// tables this service relies on if they don't already exist
func (d *Database) Init() {
	// connect to database using configuration created from environment variables
	//
	// from https://godoc.org/github.com/lib/pq
	// "Most environment variables as specified at http://www.postgresql.org/docs/current/static/libpq-envars.html
	// supported by libpq are also supported by pq."
	connector, err := pq.NewConnector("")
	if err != nil {
		panic(err)
	}

	d.db = sql.OpenDB(connector)

	// create tables
	d.createTables()
}

// Close closes connections to the database
func (d *Database) Close() error {
	return d.Close()
}

// createTables creates the tables that this service relies on
func (d *Database) createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS counter
	(
		counter_id INTEGER PRIMARY KEY,
		val INTEGER NOT NULL
	)`

	if _, err := d.db.Exec(query); err != nil {
		panic(err)
	}
}

// IncrementCounter increments the value in the counter table
func (d *Database) IncrementCounter(counterId int) error {
	query := `
	INSERT INTO counter(counter_id, val)
	VALUES($1, 0)
	ON CONFLICT(counter_id)
	DO UPDATE
	SET val = counter.val + 1`

	_, err := d.db.Exec(query, counterId)
	return err
}

// GetCounterVal gets the value of a counter with ID counterId
func (d *Database) GetCounterVal(counterId int) (int, error) {
	query := `SELECT val FROM counter WHERE counter_id = $1`

	var val int
	return val, d.db.QueryRow(query, counterId).Scan(&val)
}
