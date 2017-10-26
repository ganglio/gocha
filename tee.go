package gocha

import "fmt"

// Tee is a struct representing the channel duplicator
type Tee struct {
	out []chan interface{}
	in  chan interface{}
}

// NewTee create a new Tee object
func NewTee() *Tee {
	t := &Tee{in: make(chan interface{})}

	go func() {
		for e := range t.in {
			fmt.Printf("%#v", e)
			for _, c := range t.out {
				c <- e
			}
		}
		for _, c := range t.out {
			close(c)
		}
	}()

	return t
}

// In returns the input channel
func (t *Tee) In() chan<- interface{} {
	return t.in
}

// AddOuts adds output channels to the Tee
func (t *Tee) AddOuts(out ...chan interface{}) {
	t.out = append(t.out, out...)
}

// OutLen returns the number of output channels
func (t *Tee) OutLen() int {
	return len(t.out)
}
