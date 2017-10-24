package gocha

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestAddChannel(t *testing.T) {
	m := NewMux()
	c := make(chan interface{})
	m.AddChannel(c)
	Convey("AddChannel works as expected", t, func() {
		c <- 1
		a := <-m.Out()
		So(a, ShouldEqual, 1)
	})
}

func TestAddChannelCount(t *testing.T) {
	m := NewMux()
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	m.AddChannel(c1)
	m.AddChannel(c2)

	Convey("Count should track the number of open channels added", t, func() {
		So(m.Count(), ShouldEqual, 2)

		close(c1)
		time.Sleep(100 * time.Millisecond)
		So(m.Count(), ShouldEqual, 1)

		close(c2)
		time.Sleep(100 * time.Millisecond)
		So(m.Count(), ShouldEqual, 0)

		_, ok := <-m.Out()
		So(ok, ShouldBeFalse)
	})
}

func TestMultipleAddChannel(t *testing.T) {
	m := NewMux()
	ch := m.Out()
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	m.AddChannel(c1)
	m.AddChannel(c2)

	go func() {
		Convey("Adding multiple channel and delivering items in order", t, func() {
			cnt := 1
			for elem := range ch {
				So(elem, ShouldEqual, cnt)
				cnt += 1
			}
		})
	}()

	go func() {
		c1 <- 1
		c2 <- 2
		c2 <- 3
		c1 <- 4
	}()
}
