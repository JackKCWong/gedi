package internal

import (
	"bufio"
	"io"
)

var _ = (RecordReader)(LineReader{})

type LineReader struct{}

// Read a input line by line and return a `chan string` of lines
func (l LineReader) Read(r io.Reader) (chan Record, error) {
	var lines = make(chan Record)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	go func() {
		defer close(lines)
		i := 0
		for scanner.Scan() {
			i++
			line := scanner.Text()
			lines <- Record{
				i,
				line,
				line,
			}
		}
	}()

	return lines, nil
}

// var _ = (RecordReader)(JsonLReader{})

// type JsonLReader struct {
// 	reader LineReader
// }

// // Read implements RecordReader and reads a jsonl file
// func (j JsonLReader) Read(r io.Reader) (chan Record, error) {
// 	var out = make(chan Record)
// 	x, err := j.reader.Read(r)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read jsonl: %w", err)
// 	}

// 	go func() {
// 		defer close(out)
// 		for line := range x {
// 			var rec = make(map[string]any)
// 			err := json.Unmarshal([]byte(line.(string)), &rec)
// 			if err != nil {
// 				fmt.Printf("failed to unmarshal json: %q\n", err)
// 				continue
// 			}

// 			out <- rec
// 		}
// 	}()

// 	return out, nil
// }
