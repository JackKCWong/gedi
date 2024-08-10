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
	env := map[string]any{
		"ix": r0.lineno,
		"x": r0.parsed,
	}

	exp, err := Compile(f.Expr, env)
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)
	consume := func(r Record) {
		env["ix"] = r.lineno
		env["x"] = r.parsed
		res, err := expr.Run(exp, env)
		if err != nil {
			out <- fmt.Sprintf("error in running expr: %q", err)
		} else {
			if b, ok := res.(bool); ok && b {
				out <- r.raw
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
