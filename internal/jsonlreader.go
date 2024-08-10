package internal

import (
	"bufio"
	"fmt"
	"io"

	jsoniter "github.com/json-iterator/go"
)

var jsonit = jsoniter.ConfigCompatibleWithStandardLibrary
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
			var jsonobj = make(map[string]any)
			err := jsonit.Unmarshal([]byte(raw), &jsonobj)
			if err != nil {
				fmt.Printf("failed to unmarshal json: %q\n", err)
				continue
			}

			out <- &record{
				i,
				raw,
				jsonobj,
			}
		}
	}()

	return out, nil
}
