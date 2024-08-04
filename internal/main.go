package internal

import (
	"fmt"
	"io"
)

type RecordReader interface {
	Read(r io.Reader) (chan any, error)
}

type RecordProcessor interface {
	Process(<-chan any) (chan string, error)
}

type Gedi struct {
	reader    RecordReader
	processor RecordProcessor
}

func New(reader RecordReader, processor RecordProcessor) *Gedi {
	return &Gedi{
		reader:    reader,
		processor: processor,
	}
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
