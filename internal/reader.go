package internal

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
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
		for scanner.Scan() {
			line := scanner.Text()
			lines <- Record{
				line,
				line,
			}
		}
	}()

	return lines, nil
}

var _ = (RecordReader)(CsvReader{})

type CsvReader struct {
	reader LineReader
}

// Read a csv file and return a `chan []string` of rows and cells
func (c CsvReader) Read(r io.Reader) (chan Record, error) {
	lines, err := c.reader.Read(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}

	records := make(chan Record)
	go func() {
		defer close(records)
		buf := bytes.Buffer{}
		for l := range lines {
			buf.Reset()
			buf.WriteString(l.raw)
			csvRd := csv.NewReader(&buf)
			record, err := csvRd.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("failed to read csv: %q\n", err)
				continue
			}

			records <- Record{
				raw:    l.raw,
				parsed: record,
			}
		}
	}()

	return records, nil
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
