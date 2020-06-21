package pgsql

import (
	"context"
	"github.com/jackc/pgx"
)

type Database struct {
	db *pgx.Conn
	ctx context.Context
}

// NewDatabase creates and initializes a new Database
func NewDatabase() *Database {
	db := &Database{ctx: context.Background()}
	db.Init()

	return db
}

// Init connects to a PostgreSQL instance and creates the
// tables this service relies on if they don't already exist
func (d *Database) Init() {
	// read the PostgreSQL connection info from the environment
	config, err := pgx.ParseConfig("")
	if err != nil {
		panic(err)
	}

	// connect to database using configuration created from environment variables
	if d.db, err = pgx.ConnectConfig(d.ctx, config); err != nil {
		panic(err)
	}

	// create tables
	d.createTables()
}

// Close closes connections to the database
func (d *Database) Close() error {
	return d.db.Close(d.ctx)
}

// createTables creates the tables that this service relies on
func (d *Database) createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS counter
	(
		counter_id INTEGER PRIMARY KEY,
		val INTEGER NOT NULL
	)`

	if _, err := d.db.Exec(d.ctx, query); err != nil {
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

	_, err := d.db.Exec(d.ctx, query, counterId)
	return err
}

// GetCounterVal gets the value of a counter with ID counterId
func (d *Database) GetCounterVal(counterId int) (int, error) {
	query := `SELECT val FROM counter WHERE counter_id = $1`

	var val int
	return val, d.db.QueryRow(d.ctx, query, counterId).Scan(&val)
}
