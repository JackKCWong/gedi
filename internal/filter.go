package internal

import (
	"fmt"

	"github.com/expr-lang/expr"
)

var _ = (RecordProcessor)(Filter{})

type Filter struct {
	Expr string
}

// Process implements Processor. It prints out the original record if the expr eval to true
func (f Filter) Process(input <-chan Record) (chan string, error) {
	r0 := <-input
	exp, err := Compile(f.Expr, r0.Parsed())
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)
	consume := func(r Record) {
		res, err := expr.Run(exp, r.Parsed())
		if err != nil {
			out <- fmt.Sprintf("error in running expr: %q", err)
		} else {
			if b, ok := res.(bool); ok && b {
				out <- r.String()
			}
		}
	}

	go func() {
		defer close(out)
		consume(r0)
		for r := range input {
			consume(r)
		}
	}()

	return out, nil
}
