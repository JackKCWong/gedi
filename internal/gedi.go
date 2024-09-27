package internal

import (
	"fmt"
	"io"
	"reflect"

	"github.com/expr-lang/expr/ast"
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

type detector struct {
	isReducer bool
}

func (d *detector) Visit(node *ast.Node) {
	if _, ok := (*node).(*ast.ClosureNode); ok {
		d.isReducer = true
	}
}

func InferProcess(exp string) (RecordProcessor, error) {
	vm, err := Compile(exp, map[string]any{})
	if err != nil {
		return nil, fmt.Errorf("failed to compile expr: %w", err)
	}

	root := vm.Node()
	d := &detector{}
	ast.Walk(&root, d)
	if d.isReducer {
		return Reducer{Expr: exp}, nil
	}

	if root.Type().AssignableTo(reflect.TypeOf(true)) {
		return Filter{Expr: exp}, nil
	}

	return Mapper{Expr: exp}, nil
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
func (r *record) Parsed() map[string]any {
	return r.parsed
}

// String implements Record.
func (r *record) String() string {
	return r.raw
}
