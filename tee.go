package gocha

import "fmt"

// Tee is a struct representing the channel duplicator
type Tee struct {
	out []chan interface{}
	in  chan interface{}
}

// Create a new Tee object
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

// The input channel
func (t *Tee) In() chan<- interface{} {
	return t.in
}

// Adds outs to the tee
func (t *Tee) AddOuts(out ...chan interface{}) {
	t.out = append(t.out, out...)
}

// Returns the number of outs attached to the tee
func (t *Tee) OutLen() int {
	return len(t.out)
}
