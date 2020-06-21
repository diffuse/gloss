package gloss

// Database represents a wrapper around a connection to a database
//
// additional methods can be added here, then implemented in sub-packages
type Database interface {
	Init()
	Close() error

	// these are examples of business logic for the
	// counter service, replace them with your own
	IncrementCounter(counterId int) error
	GetCounterVal(counterId int) (int, error)
}
