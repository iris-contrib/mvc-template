package services

import (
	"math"
	"sync/atomic"
)

// GlobalCounter is a Counter which
// atomically increments until Max.
//
// This is one of the safest implementations for global counting.
// Test it: wrk -c 20 -d 3s http://localhost:8080/api/counter/increment
type GlobalCounter struct {
	value uint64
	Max   uint64
}

// This is a compile-time check,
// we make sure that GlobalCounter completes the Counter interface.
var _ Counter = (*GlobalCounter)(nil)

// NewGlobalCounter returns a global counter.
func NewGlobalCounter() *GlobalCounter {
	return &GlobalCounter{Max: math.MaxUint64}
}

// Increment increments the Value.
// The value cannot exceed the Max one.
// It uses Compare and Swap with the atomic package.
//
// Returns the new number value.
func (c *GlobalCounter) Increment() (newValue uint64) {
	for {
		prev := atomic.LoadUint64(&c.value)
		newValue = prev + 1

		if newValue >= c.Max {
			newValue = 0
		}

		if atomic.CompareAndSwapUint64(&c.value, prev, newValue) {
			break
		}
	}

	return
}

// Get returns the current counter without incrementing.
func (c *GlobalCounter) Get() uint64 {
	return atomic.LoadUint64(&c.value)
}
