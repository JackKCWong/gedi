package internal

import (
	"fmt"

	csp "github.com/JackKCWong/chansport"
	"github.com/expr-lang/expr"
)

var _ = (RecordProcessor)(Filter{})

type Filter struct {
	Expr     string
	Parallel int
}

// Process implements Processor. It prints out the original record if the expr eval to true
func (f Filter) Process(input <-chan Record) (chan string, error) {
	r0 := <-input
	env := map[string]any{
		"ix": r0.LineNo(),
		"x":  r0.Parsed(),
	}

	exp, err := Compile(f.Expr, env)
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	consume := func(r Record) string {
		env := map[string]any{
			"ix": r.LineNo(),
			"x":  r.Parsed(),
		}
		res, err := expr.Run(exp, env)
		if err != nil {
			return fmt.Sprintf("error in running expr: %q", err)
		}

		if b, ok := res.(bool); ok && b {
			return r.Raw()
		}

		return ""
	}

	out := make(chan string, 1)
	o0 := consume(r0)
	if o0 != "" {
		out <- o0
	}

	if f.Parallel > 1 {
		go func() {
			defer close(out)
			for o := range csp.MapParallel(input, consume, f.Parallel) {
				if o != "" {
					out <- o
				}
			}
		}()
	} else {
		go func() {
			defer close(out)
			for i := range input {
				o := consume(i)
				if o != "" {
					out <- o
				}
			}
		}()
	}

	return out, nil
}
