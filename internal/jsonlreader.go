package internal

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ = (RecordReader)(JsonLReader{})

type JsonLReader struct {
	reader LineReader
}

// Read implements RecordReader and reads a jsonl file
func (j JsonLReader) Read(r io.Reader) (chan Record, error) {
	var out = make(chan Record)
	x, err := j.reader.Read(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read jsonl: %w", err)
	}

	go func() {
		defer close(out)
		for r := range x {
			var jsonobj = make(map[string]any)
			err := json.Unmarshal([]byte(r.raw), &jsonobj)
			if err != nil {
				fmt.Printf("failed to unmarshal json: %q\n", err)
				continue
			}

			r.parsed = jsonobj
			out <- r
		}
	}()

	return out, nil
}
