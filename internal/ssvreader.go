package internal

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

// SsvReader reads records from a space separated values file
type SsvReader struct {
	MaxNumOfFields int
}

var _ = (RecordReader)(&SsvReader{})
var re = regexp.MustCompile(`\s+`)

func (r *SsvReader) Read(in io.Reader) (chan Record, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	var out = make(chan Record)
	go func() {
		defer close(out)
		lineno := 0
		for scanner.Scan() {
			lineno++

			raw := scanner.Text()
			line := strings.TrimSpace(raw)
			if len(line) == 0 {
				continue
			}

			fields := re.Split(line, r.MaxNumOfFields)
			if len(fields) == 0 {
				continue
			}

			out <- &record{lineno: lineno, raw: raw, parsed: map[string]any{"x": fields}}
		}
	}()

	return out, nil
}