package gocha

type selector func(interface{}) int

// Switch is the struct representing the channel switch
type Switch struct {
	in  chan interface{}
	out []chan interface{}
	sel selector
}

// NewSwitch creates a new Switch object
func NewSwitch(s selector) *Switch {
	sw := &Switch{in: make(chan interface{}), sel: s}
	go func() {
		for e := range sw.in {
			if len(sw.out) > 0 {
				i := sw.sel(e) % len(sw.out)
				sw.out[i] <- e
			}
		}
		for _, c := range sw.out {
			close(c)
		}
	}()
	return sw
}

// AddOuts adds output channels to the switch
func (s *Switch) AddOuts(ch ...chan interface{}) {
	s.out = append(s.out, ch...)
}

// OutLen returns the number of output channels
func (s *Switch) OutLen() int {
	return len(s.out)
}

// In returns the input channel
func (s *Switch) In() chan<- interface{} {
	return s.in
}
