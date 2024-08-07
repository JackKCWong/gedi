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

const BUF_SIZE = 4096

// Read a csv file and return a `chan []string` of rows and cells
func (c *CsvReader) Read(r io.Reader) (chan Record, error) {
	rawBuf := bytes.Buffer{}
	records := make(chan Record)
	csvRd := csv.NewReader(io.TeeReader(r, &rawBuf))

	go func() {
		defer close(records)
		var lastOffset int64
		var shift int64
		for {
			record, err := csvRd.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("failed to read csv: %q\n", err)
				continue
			}

			raw := rawBuf.String()[lastOffset:csvRd.InputOffset()-shift]
			records <- Record{
				raw:    strings.TrimRight(raw, "\r\n"),
				parsed: record,
			}

			if lastOffset > BUF_SIZE {
				remain := rawBuf.String()[csvRd.InputOffset()-shift:]
				rawBuf.Reset()
				rawBuf.WriteString(remain)
				shift = csvRd.InputOffset()
				lastOffset = 0
			} else {
				lastOffset = csvRd.InputOffset() - shift
			}
		}
	}()

	return records, nil
}
