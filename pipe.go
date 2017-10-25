package gocha

// Proc is a type abstracting a pipe processor
type Proc func(interface{}) interface{}

// Pipe is a struct representing the processed pipe
type Pipe struct {
	in    chan interface{}
	out   chan interface{}
	steps []Proc
}

// Create a new Pipe object
func NewPipe() *Pipe {
	p := &Pipe{
		in:  make(chan interface{}),
		out: make(chan interface{}),
	}
	go func() {
		for e := range p.in {
			for _, pr := range p.steps {
				e = pr(e)
			}
			p.out <- e
		}
		close(p.out)
	}()
	return p
}

// The input channel
func (p *Pipe) In() chan<- interface{} {
	return p.in
}

// The Output channel
func (p *Pipe) Out() <-chan interface{} {
	return p.out
}

// Adds procs to the pipe
func (p *Pipe) AddProcs(pr ...Proc) {
	p.steps = append(p.steps, pr...)
}

// Returns the numbe of procs applied to the pipe
func (p *Pipe) ProcsLen() int {
	return len(p.steps)
}
