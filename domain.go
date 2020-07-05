package gloss

import "net/http"

// Database represents a wrapper around a database connection + driver
//
// additional methods can be added here, then implemented in custom packages
type Database interface {
	// Init connects to the database, creates tables, etc...
	Init()

	// Close shuts down connections to the database
	Close() error

	// these are examples of business logic for the
	// counter service, replace them with your own
	IncrementCounter(counterId int) error
	GetCounterVal(counterId int) (int, error)
}

// Router represents a router that implements
// both the net/http Handler interface, as well as
// a method to associate a database with its handlers
type Router interface {
	http.Handler
	SetDb(db Database) error
}
