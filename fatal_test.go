package errors

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIsFatal(t *testing.T) {

	Convey("Given a no error", t, func() {
		var err error
		Convey("When calling IsFatal", func() {
			f := IsFatal(err)
			Convey("The result should be false", func() {
				So(f, ShouldBeFalse)
			})
		})
	})

	Convey("Given an error that does not implement Fatal", t, func() {
		err := fmt.Errorf("I'm an error!")
		Convey("When calling IsFatal", func() {
			f := IsFatal(err)
			Convey("The result should be false", func() {
				So(f, ShouldBeFalse)
			})

		})
	})

	Convey("Given a non-fatal error that implements Fatal", t, func() {
		err := fatal{message: "I'm a non-fatal, fatal error", isfatal: false}
		Convey("When calling IsFatal", func() {
			f := IsFatal(err)
			Convey("The result should be false", func() {
				So(f, ShouldBeFalse)
			})

		})
	})

	Convey("Given a fatal error that implements Fatal", t, func() {
		err := fatal{message: "I'm a fatal, fatal error", isfatal: true}
		Convey("When calling IsFatal", func() {
			f := IsFatal(err)
			Convey("The result should be true", func() {
				So(f, ShouldBeTrue)
			})
		})

	})

}

func TestFatal(t *testing.T) {

	Convey("Given a fatal error", t, func() {
		err := fatal{message: "I'm a fatal, fatal error", isfatal: false}

		Convey("The error should implement Fatal", func() {
			_, ok := error(err).(Fatal)
			So(ok, ShouldBeTrue)
		})

		Convey("When isfatal is set to false", func() {
			err.isfatal = false
			Convey("IsFatal() should return false", func() {
				So(err.IsFatal(), ShouldBeFalse)
			})
		})

		Convey("When isfatal is set to true", func() {
			err.isfatal = true
			Convey("IsFatal() should return true", func() {
				So(err.IsFatal(), ShouldBeTrue)
			})
		})

		Convey("When calling Error()", func() {
			err.message = "foo"
			Convey("The message should be correct", func() {
				So(err.Error(), ShouldEqual, "foo")
			})
		})

	})

}

func TestFatalf(t *testing.T) {

	Convey("Given a new fatal error", t, func() {
		err := Fatalf("I'm a fatal error!")
		Convey("The error should implement Fatal", func() {
			_, ok := err.(Fatal)
			So(ok, ShouldBeTrue)
		})

		Convey("When IsFatal is called", func() {
			f := err.(Fatal).IsFatal()
			Convey("The result should be true", func() {
				So(f, ShouldBeTrue)
			})
		})

		Convey("When calling Error()", func() {
			Convey("The message should be correct", func() {
				So(err.Error(), ShouldEqual, "I'm a fatal error!")
			})
		})

	})

}
