package internal

import (
	"bytes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadLines(t *testing.T) {
	Convey("it can read files line by line", t, func() {
		g := Gedi{
			reader: LineReader{},
			processor: Filter{
				Expr: "atoi(x) > 10",
			},
		}

		err := g.Run(bytes.NewBufferString(
			"1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n11\n12\n13\n14\n15\n16\n17\n18\n19\n20\n",
		))

		So(err, ShouldBeNil)
	})
}
