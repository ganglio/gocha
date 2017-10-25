package gocha

import (
	"sync"
)

// Mux is a struct representing the muxed channels
type Mux struct {
	ch chan interface{}
	sync.RWMutex
	count int
}

// create a new Mux object
func NewMux() *Mux {
	m := &Mux{ch: make(chan interface{}), count: 0}
	return m
}

// Returns the number of muxed channels
func (m *Mux) Count() int {
	m.RLock()
	defer m.RUnlock()
	return m.count
}

func (m *Mux) Out() <-chan interface{} {
	return m.ch
}

// Adds a channel to the mux
// If the channel gets closed the goroutine handling it completes and the counter decreases
// If the counter reaches 0 the muxed channel automatically closes
func (m *Mux) addChannel(c chan interface{}) {
	m.Lock()
	defer m.Unlock()
	m.count += 1

	go func(c chan interface{}) {
		for elem := range c {
			m.ch <- elem
		}
		m.Lock()
		defer m.Unlock()
		m.count -= 1

		if m.count == 0 {
			close(m.ch)
		}
	}(c)
}

// Add an bunch of channels to the Mux
func (m *Mux) AddChannels(c ...chan interface{}) {
	for _, ch := range c {
		m.addChannel(ch)
	}
}
