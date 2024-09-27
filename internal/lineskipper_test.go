package internal

import (
	"bufio"
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLineSkipper(t *testing.T) {
	Convey("Given a LineSkipper with 8 lines to skip", t, func() {
		buf := bytes.NewBufferString("1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n")
		skipper := LineSkipper{NumOfLines: 8, Reader: buf}

		scanner := bufio.NewScanner(&skipper)

		So(scanner.Scan(), ShouldBeTrue)
		So(scanner.Text(), ShouldEqual, "9")
		So(scanner.Scan(), ShouldBeTrue)
		So(scanner.Text(), ShouldEqual, "10")
		So(scanner.Scan(), ShouldBeFalse)
	})
}
