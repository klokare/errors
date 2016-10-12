package errors

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMultiple(t *testing.T) {

	Convey("Given Multiple errors", t, func() {
		m := new(Multiple)

		Convey("When there are zero errors", func() {
			Convey("IsFatal should be false", func() {
				So(m.IsFatal(), ShouldBeFalse)
			})
		})

		Convey("When there is one or more errors but all are non-fatal", func() {
			m.Add(fmt.Errorf("This is non fatal error 1"))
			m.Add(fmt.Errorf("This is non fatal error 2"))
			Convey("IsFatal should be false", func() {
				So(m.IsFatal(), ShouldBeFalse)
			})
		})

		Convey("When there is one or more errors and at least one is fatal", func() {
			m.Add(fmt.Errorf("This is non fatal error 1"))
			m.Add(fmt.Errorf("This is non fatal error 2"))
			m.Add(Fatalf("This is a fatal error, sucka!"))
			Convey("IsFatal should be true", func() {
				So(m.IsFatal(), ShouldBeTrue)
			})
		})

	})

}
