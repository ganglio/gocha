package gocha

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAddProcSingle(t *testing.T) {
	Convey("AddProc works as expected", t, func() {
		p := NewPipe()
		Convey("When freshly instantiated ProcsLen is zero", func() {
			So(p.ProcsLen(), ShouldEqual, 0)
		})
		p.AddProcs(func(i interface{}) interface{} {
			return i
		})
		Convey("Adding a proc increses the ProcsLen to 1", func() {
			So(p.ProcsLen(), ShouldEqual, 1)
		})
	})
}

func TestAddProcComposing(t *testing.T) {
	p := NewPipe()
	p.AddProcs(func(i interface{}) interface{} {
		return 2 + i.(int)
	})

	Convey("Testing that a proc is applied correctly", t, func() {
		p.In() <- 1
		So(<-p.Out(), ShouldEqual, 3)
	})

	p.AddProcs(func(i interface{}) interface{} {
		return i.(int)/3 + 1
	})

	Convey("Adding a new proc compose correctly", t, func() {
		p.In() <- 1
		So(<-p.Out(), ShouldEqual, 2)
	})
}

func TestAddProcs(t *testing.T) {
	p1 := func(i interface{}) interface{} {
		return 3*i.(int) + 1
	}
	p2 := func(i interface{}) interface{} {
		return i.(int) / 2
	}
	p3 := func(i interface{}) interface{} {
		return i.(int) * i.(int)
	}

	Convey("Adding multiple procs at the same time depends on the addition order", t, func() {
		p := NewPipe()
		p.AddProcs(p1, p2, p3)
		p.In() <- 1
		So(<-p.Out(), ShouldEqual, 4)

		p = NewPipe()
		p.AddProcs(p2, p1, p3)
		p.In() <- 2
		So(<-p.Out(), ShouldEqual, 16)
	})

	Convey("Closing the input closes the output", t, func() {
		p := NewPipe()
		close(p.In())

		_, ok := <-p.Out()
		So(ok, ShouldBeFalse)
	})
}
