package internal

import (
	"fmt"
	"io"
)

type InputReader interface {
	Read(r io.Reader) (chan any, error)
}

type Processor interface {
	Process(<-chan any) (chan string, error)
}

type Gedi struct {
	reader    InputReader
	processor Processor
}

func (g Gedi) Run(input io.Reader) error {
	x, err := g.reader.Read(input)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	y, err := g.processor.Process(x)
	if err != nil {
		return fmt.Errorf("failed to process input: %w", err)
	}

	for o := range y {
		fmt.Println(o)
	}

	return nil
}
