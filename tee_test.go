package gocha

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOutCount(t *testing.T) {
	tee := NewTee()
	Convey("Upon creation OutLen is zero", t, func() {
		So(tee.OutLen(), ShouldEqual, 0)
	})

	Convey("Adding an out OutLen increases to 1", t, func() {
		c := make(chan interface{})
		tee.AddOuts(c)
		So(tee.OutLen(), ShouldEqual, 1)
	})
}

func TestAddOut(t *testing.T) {
	tee := NewTee()
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	tee.AddOuts(c1, c2)
	Convey("A message in is duplicated to all the outs", t, func() {
		tee.In() <- 1
		So(<-c1, ShouldEqual, 1)
		So(<-c2, ShouldEqual, 1)
	})
	Convey("Closing in closes all the outs", t, func() {
		close(tee.In())

		_, ok := <-c1
		So(ok, ShouldBeFalse)

		_, ok = <-c2
		So(ok, ShouldBeFalse)
	})
}
