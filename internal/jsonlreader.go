package internal

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

var _ = (RecordReader)(&JsonLReader{})

type JsonLReader struct{}

// Read implements RecordReader and reads a jsonl file
func (j *JsonLReader) Read(r io.Reader) (chan Record, error) {
	var out = make(chan Record)
	go func() {
		defer close(out)
		scanner := bufio.NewScanner(r)
		scanner.Split(bufio.ScanLines)
		i := 0
		for scanner.Scan() {
			i++
			raw := scanner.Text()
			var obj = make(map[string]any)
			err := json.Unmarshal([]byte(raw), &obj)
			if err != nil {
				fmt.Printf("failed to unmarshal json: %q\n", err)
				continue
			}

			out <- &record{
				i,
				raw,
				map[string]any{
					"ix": i,
					"x":  obj,
				},
			}
		}
	}()

	return out, nil
}
