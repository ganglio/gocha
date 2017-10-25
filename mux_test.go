package gocha

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestAddChannel(t *testing.T) {
	m := NewMux()
	c := make(chan interface{})
	m.AddChannels(c)

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

	m.AddChannels(c1, c2)

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
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	m.AddChannels(c1, c2)

	go func() {
		defer close(c1)
		defer close(c2)
		c1 <- true
		c2 <- true
		c2 <- true
		c1 <- true
	}()

	Convey("Adding multiple channel and delivering items in order", t, func() {
		cnt := 0
		for elem := range m.Out() {
			So(elem, ShouldBeTrue)
			cnt += 1
		}
		So(cnt, ShouldEqual, 4)
	})
}
