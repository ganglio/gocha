package gocha

import (
	"sync"
)

type Mux struct {
	mu    sync.RWMutex
	count int
	Out   chan interface{}
}

func NewMux() *Mux {
	return &Mux{Out: make(chan interface{})}
}

func (m *Mux) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.count
}

func (m *Mux) AddChannel(c chan interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.count += 1

	go func(c chan interface{}) {
		for elem := range c {
			m.Out <- elem
		}
		m.mu.Lock()
		defer m.mu.Unlock()
		m.count -= 1

		if m.count == 0 {
			close(m.Out)
		}
	}(c)
}
