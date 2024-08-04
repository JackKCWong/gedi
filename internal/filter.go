package internal

import (
	"fmt"

	"github.com/expr-lang/expr"
)

var _ = (Processor)(Filter{})

type Filter struct {
	expr string
}

// Process implements Processor. It prints out the original record if the expr eval to true
func (f Filter) Process(input <-chan any) (chan string, error) {
	env := map[string]any{
		"x": "",
	}

	exp, err := Compile(f.expr, env)
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)

	go func() {
		defer close(out)
		for x := range input {
			env["x"] = x
			res, err := expr.Run(exp, env)
			if err != nil {
				out <- fmt.Sprintf("error in running expr: %q\n", err)
				continue
			}

			if b, ok := res.(bool); ok && b {
				out <- fmt.Sprintf("%s", x)
			}
		}
	}()

	return out, nil
}
