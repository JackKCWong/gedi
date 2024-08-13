package internal

import (
	"encoding/json"
	"fmt"

	"github.com/expr-lang/expr"
)

var _ = (RecordProcessor)(Mapper{})

type Mapper struct {
	Expr string
}

// Process implements Processor. It prints out the result of the Expr
func (m Mapper) Process(input <-chan Record) (chan string, error) {
	r0 := <-input
	env := map[string]any{
		"ix": r0.LineNo(),
		"x":  r0.Parsed(),
	}

	exp, err := Compile(m.Expr, env)
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)
	consume := func(r Record) {
		env["ix"] = r.LineNo()
		env["x"] = r.Parsed()
		result, err := expr.Run(exp, env)
		if err != nil {
			out <- fmt.Sprintf("error in running expr: %q", err)
		} else {
			if j, ok := result.(map[string]interface{}); ok {
				jsonstr, err := json.Marshal(j)
				if err != nil {
					out <- fmt.Sprintf("failed to marshal to json: %q", err)
				} else {
					out <- string(jsonstr)
				}
			} else {
				out <- fmt.Sprintf("%v", result)
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
