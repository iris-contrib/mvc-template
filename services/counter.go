package services

// Counter is the interface which
// different implementations of counters should implement.
type Counter interface {
	Increment() uint64
	Get() uint64
}
