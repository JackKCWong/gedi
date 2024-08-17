package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

	go func() {
		defer close(records)
		i := 0
		kx := ""
		for {
			for {
				// read until we find a '['
				t, err := dec.Token()
				if err == io.EOF {
					return
				}
				if err != nil {
					fmt.Fprintln(os.Stderr, fmt.Errorf("failed to read json: %w", err))
					return
				}
				if name, ok := t.(string); ok {
					kx = name
				}
				if delim, ok := t.(json.Delim); ok && delim == '[' {
					break
				}
			}

			for dec.More() {
				// read array elements
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
						"kx": kx,
					},
				}
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
	str, err := json.Marshal(j.parsed["x"])
	if err != nil {
		panic(err)
	}

	return string(str)
}
