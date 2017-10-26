package gocha

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewSwitch(t *testing.T) {
	Convey("Initialisation works as expected", t, func() {
		sw := NewSwitch(func(i interface{}) int {
			return i.(int)
		})
		Convey("When freshly instantiated OutLen is zero", func() {
			So(sw.OutLen(), ShouldEqual, 0)
		})
		Convey("When freshly instantiated sending is a noop", func() {
			sw.In() <- 1
			sw.In() <- 1
			So(sw.OutLen(), ShouldEqual, 0)
		})
	})
}

func TestAddOuts(t *testing.T) {
	Convey("AddOuts works as expected", t, func() {
		sw := NewSwitch(func(i interface{}) int {
			return i.(int)
		})
		c1 := make(chan interface{})
		sw.AddOuts(c1)
		Convey("With only one out the Switch acts as a pipe", func() {
			sw.In() <- 33
			So(<-c1, ShouldEqual, 33)
		})
		c2 := make(chan interface{})
		sw.AddOuts(c2)
		Convey("With two outs the Switch routes correctly", func() {
			sw.In() <- 33
			So(<-c2, ShouldEqual, 33)
			sw.In() <- 34
			So(<-c1, ShouldEqual, 34)
		})
	})
}

func TestClosingSwitch(t *testing.T) {
	Convey("Closing the input channel closes the outs", t, func() {
		sw := NewSwitch(func(i interface{}) int {
			return i.(int)
		})
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		sw.AddOuts(c1, c2)

		close(sw.In())
		_, ok1 := <-c1
		_, ok2 := <-c2
		So(ok1, ShouldBeFalse)
		So(ok2, ShouldBeFalse)
	})
}
