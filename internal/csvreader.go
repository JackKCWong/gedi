package internal

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

var _ = (RecordReader)((*CsvReader)(nil))

type CsvReader struct {
}

// Read a csv file and return a `chan []string` of rows and cells
func (c *CsvReader) Read(r io.Reader) (chan Record, error) {
	rawBuf := bytes.Buffer{}
	records := make(chan Record)
	csvRd := csv.NewReader(io.TeeReader(r, &rawBuf))

	go func() {
		defer close(records)
		var lastOffset int64
		for {
			record, err := csvRd.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("failed to read csv: %q\n", err)
				continue
			}

			records <- Record{
				raw:    strings.TrimRight(rawBuf.String()[lastOffset:csvRd.InputOffset()], "\r\n"),
				parsed: record,
			}

			lastOffset = csvRd.InputOffset()
		}
	}()

	return records, nil
}
