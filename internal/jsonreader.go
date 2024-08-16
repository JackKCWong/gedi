package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

type JsonReader struct{}
type jsonrecord struct {
	lineno int
	parsed map[string]any
}

var _ = (RecordReader)(&JsonReader{})
var _ = (Record)(&jsonrecord{})

// Read implements RecordReader.
func (j *JsonReader) Read(r io.Reader) (chan Record, error) {
	records := make(chan Record)
	dec := json.NewDecoder(r)
	t, err := dec.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to read json array: %w", err)
	}

	if t.(json.Delim) != '[' {
		return nil, fmt.Errorf("expected json array, got %q", t)
	}

	go func() {
		defer close(records)
		i := 0
		for dec.More() {
			i++
			var obj map[string]any
			err := dec.Decode(&obj)
			if err != nil {
				fmt.Printf("failed to decode json: %q\n", err)
				continue
			}
			records <- &jsonrecord{
				lineno: i,
				parsed: map[string]any{
					"ix": i,
					"x":  obj,
				},
			}
		}
	}()

	return records, nil
}

// LineNo implements Record.
func (j *jsonrecord) LineNo() int {
	return j.lineno
}

// Parsed implements Record.
func (j *jsonrecord) Parsed() map[string]any {
	return j.parsed
}

// String implements Record.
func (j *jsonrecord) String() string {
	str, err := json.Marshal(j.parsed)
	if err != nil {
		panic(err)
	}

	return string(str)
}
