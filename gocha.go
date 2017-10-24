package gocha

type Mux struct {
	Count int
	Out   chan interface{}
}

func NewMux() *Mux {
	return &Mux{Out: make(chan interface{})}
}

func (m *Mux) AddChannel(c chan interface{}) {
	m.Count += 1
	go func(c chan interface{}) {
		for elem := range c {
			m.Out <- elem
		}
		m.Count -= 1
		if m.Count == 0 {
			close(m.Out)
		}
	}(c)
}
