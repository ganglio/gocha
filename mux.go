package gocha

import (
	"sync"
)

// Mux is a struct representing the muxed channels
type Mux struct {
	out chan interface{}
	sync.RWMutex
	count int
}

// NewMux creates a new Mux object
func NewMux() *Mux {
	m := &Mux{out: make(chan interface{}), count: 0}
	return m
}

// Count returns the number of muxed channels
func (m *Mux) Count() int {
	m.RLock()
	defer m.RUnlock()
	return m.count
}

// Out returns the output channel
func (m *Mux) Out() <-chan interface{} {
	return m.out
}

func (m *Mux) addChannel(c chan interface{}) {
	m.Lock()
	defer m.Unlock()
	m.count++

	go func(c chan interface{}) {
		for elem := range c {
			m.out <- elem
		}
		m.Lock()
		defer m.Unlock()
		m.count--

		if m.count == 0 {
			close(m.out)
		}
	}(c)
}

// AddChannels adds channels (sic!) to the mux
// If the channel gets closed the goroutine handling it completes and the counter decreases
// If the counter reaches 0 the muxed channel automatically closes
func (m *Mux) AddChannels(c ...chan interface{}) {
	for _, ch := range c {
		m.addChannel(ch)
	}
}
