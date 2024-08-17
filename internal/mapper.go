package internal

import (
	"encoding/json"
	"fmt"
)

var _ = (RecordProcessor)(Mapper{})

type Mapper struct {
	Expr string
}

// Process implements Processor. It prints out the result of the Expr
func (m Mapper) Process(input <-chan Record) (chan string, error) {
	r0 := <-input
	exp, err := Compile(m.Expr, r0.Parsed())
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	out := make(chan string)
	consume := func(r Record) {
		result, err := RunExpr(exp, r.Parsed())
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
