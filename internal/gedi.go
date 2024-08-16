package internal

import (
	"fmt"
	"io"
)

type Record interface {
	LineNo() int
	String() string
	Parsed() map[string]any
}

type RecordReader interface {
	Read(r io.Reader) (chan Record, error)
}

type RecordProcessor interface {
	Process(<-chan Record) (chan string, error)
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

var _ = (Record)(&record{})

type record struct {
	lineno int
	raw    string
	parsed map[string]any
}

// LineNo implements Record.
func (r *record) LineNo() int {
	return r.lineno
}

// Parsed implements Record.
func (r *record) Parsed() map[string]any{
	return r.parsed
}

// String implements Record.
func (r *record) String() string {
	return r.raw
}
