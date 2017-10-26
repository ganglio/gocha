package gocha

// Proc is a type abstracting a pipe processor
type proc func(interface{}) interface{}

// Pipe is a struct representing the processed pipe
type Pipe struct {
	in    chan interface{}
	out   chan interface{}
	steps []proc
}

// NewPipe creates a new Pipe object
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

// In returns the input channel
func (p *Pipe) In() chan<- interface{} {
	return p.in
}

// Out returns the output channel
func (p *Pipe) Out() <-chan interface{} {
	return p.out
}

// AddProcs adds processors to the pipeline
func (p *Pipe) AddProcs(pr ...proc) {
	p.steps = append(p.steps, pr...)
}

// ProcsLen eturns the number of procs applied to the pipe
func (p *Pipe) ProcsLen() int {
	return len(p.steps)
}
