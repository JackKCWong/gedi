package internal

import (
	"fmt"

	csp "github.com/JackKCWong/chansport"
	"github.com/expr-lang/expr"
)

var _ = (RecordProcessor)(Reducer{})

type Reducer struct {
	Expr string
}

// Process implements Processor. It prints out the original record if the expr eval to true
func (r Reducer) Process(input <-chan Record) (chan string, error) {
	out := make(chan string)
	go func() {
		defer close(out)
		x := csp.Reduce(input, func(agg []any, r Record) []any{
			return append(agg, r.Parsed())
		})

		env := map[string]any{
			"x": x,
		}

		exp, err := Compile(r.Expr, env)
		if err != nil {
			out <- fmt.Sprintf("failed to compile expr: %q", err)
			return
		}

		res, err := expr.Run(exp, env)
		if err != nil {
			out <- fmt.Sprintf("error in running expr: %q", err)
		} else {
			out <- fmt.Sprintf("%v", res)
		}
	}()

	return out, nil
}
