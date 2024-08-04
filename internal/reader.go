package internal

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
)

var _ = (RecordReader)(LineReader{})

type LineReader struct{}

// Read a input line by line and return a `chan string` of lines
func (l LineReader) Read(r io.Reader) (chan any, error) {
	var lines = make(chan any)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	go func() {
		defer close(lines)
		for scanner.Scan() {
			lines <- scanner.Text()
		}
	}()

	return lines, nil
}

var _ = (RecordReader)(CsvReader{})

type CsvReader struct{}

// Read a csv file and return a `chan []string` of rows and cells
func (c CsvReader) Read(r io.Reader) (chan any, error) {
	csvRd := csv.NewReader(r)
	var records = make(chan any)
	go func() {
		defer close(records)
		for {
			record, err := csvRd.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("failed to read csv: %q\n", err)
				continue
			}

			records <- record
		}
	}()

	return records, nil
}
