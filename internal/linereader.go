package internal

import (
	"bufio"
	"io"
)

var _ = (RecordReader)(&LineReader{})

type LineReader struct{}

// Read a input line by line and return a `chan string` of lines
func (l *LineReader) Read(r io.Reader) (chan Record, error) {
	var lines = make(chan Record)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	go func() {
		defer close(lines)
		i := 0
		for scanner.Scan() {
			i++
			line := scanner.Text()
			lines <- &record{
				i,
				line,
				line,
			}
		}
	}()

	return lines, nil
}
