package internal

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFilter(t *testing.T) {
	Convey("filter can print out lines that matches the expr", t, func() {
		f := Filter{
			Expr: "atoi(x) % 2 == 0",
			Parallel: 4,
		}

		x := make(chan Record)
		go func() {
			defer close(x)
			for i := 0; i < 10; i++ {
				x <- &record{
					lineno: i,
					raw:    strconv.Itoa(i),
					parsed: strconv.Itoa(i),
				}
			}
		}()

		y, err := f.Process(x)

		So(err, ShouldBeNil)
		So(next(y), ShouldEqual, "0")
		So(next(y), ShouldEqual, "2")
		So(next(y), ShouldEqual, "4")
		So(next(y), ShouldEqual, "6")
		So(next(y), ShouldEqual, "8")
		So(next(y), ShouldBeError)
	})
}

func next[T any](c <-chan T) any {
	select {
	case v, ok := <-c:
		if ok {
			return v
		} else {
			return fmt.Errorf("ended")
		}
		// timeout in 100ms
	case <-time.After(100 * time.Millisecond):
		return fmt.Errorf("timeout")
	}
}
