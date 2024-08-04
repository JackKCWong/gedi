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
	env := map[string]any{
		"x": "",
	}

	exp, err := Compile(f.Expr, env)
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)

	go func() {
		defer close(out)
		for r := range input {
			env["x"] = r.parsed
			res, err := expr.Run(exp, env)
			if err != nil {
				out <- fmt.Sprintf("error in running expr: %q", err)
				continue
			}

			if b, ok := res.(bool); ok && b {
				out <- fmt.Sprintf("%s", r.raw)
			}
		}
	}()

	return out, nil
}
