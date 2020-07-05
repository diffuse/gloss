package gloss

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
