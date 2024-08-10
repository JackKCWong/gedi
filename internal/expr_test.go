package internal

import (
	"testing"

	"github.com/araddon/dateparse"
	"github.com/expr-lang/expr"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWithin(t *testing.T) {
	Convey("within equals to onOrAfter", t, func() {
		now = dateparse.MustParse("2024-08-10")
		env := make(map[string]any)
		p, err := Compile(`date("2024-08-09") | within(-1 * day)`, env)
		So(err, ShouldBeNil)

		r, err := expr.Run(p, env)
		So(err, ShouldBeNil)
		So(r, ShouldBeTrue)
	})
}
