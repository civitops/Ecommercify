package pkg

// Error represents a handler error. It provides methods for a HTTP status
// code and embeds the built-in error interface.
type Error interface {
	error
	Status() int
}

// UserErr represents an error with an associated HTTP status code.
type UserErr struct {
	Code int
	Err  error
}

// Error allows UserErr to satisfy the error interface.
func (e UserErr) Error() string {
	return e.Err.Error()
}

// Status returns the HTTP status code.
func (e UserErr) Status() int {
	return e.Code
}
