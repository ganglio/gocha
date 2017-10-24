package gocha

import (
	"sync"
)

// Mux is a struct representing the muxed
type Mux struct {
  chan interface{}
	mu    sync.RWMutex
	count int
}

// cre
func NewMux() *Mux {
	return &Mux{Out: make(chan interface{})}
}

// Returns the number of muxed channels
func (m *Mux) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.count
}

// Adds a channel to the mux
// If the channel gets closed the goroutine handling it completes and the counter decreases
// If the counter reaches 0 the muxed channel automatically closes
func (m *Mux) AddChannel(c chan interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count += 1

	go func(c chan interface{}) {
		for elem := range c {
			m <- elem
		}
		m.mu.Lock()
		defer m.mu.Unlock()
		m.count -= 1

		if m.count == 0 {
			close(m.Out)
		}
	}(c)
}
