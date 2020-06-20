package gloss

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"strconv"
)

var (
	db       *sql.DB
	host     string
	port     uint16
	user     string
	password string
	dbname   string
	sslMode  string
)

// InitDb connects to a PostgreSQL instance and creates the
// tables this service relies on if they don't already exist
func InitDb() {
	// get connection info from environment
	host = os.Getenv("POSTGRES_HOST")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASS")
	dbname = os.Getenv("POSTGRES_DBNAME")
	sslMode = os.Getenv("POSTGRES_SSL_MODE")

	// parse the port
	portVal, err := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 16)
	if err != nil {
		panic(err)
	}
	port = uint16(portVal)

	// construct a connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslMode)

	// open connection to the database
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// check the connection
	if err := db.Ping(); err != nil {
		panic(err)
	}

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