package internal

import (
	"bufio"
	"io"
	"sync"
)

var _ = (io.Reader)(&LineSkipper{})

// LineSkipper is an io.Reader that skips the first n lines of the input
type LineSkipper struct {
	NumOfLines int
	Reader     io.Reader
	skip sync.Once
	remain io.Reader
}

// Read implements io.Reader.
func (s *LineSkipper) Read(p []byte) (n int, err error) {
	s.skip.Do(func() {
		scanner := bufio.NewScanner(s.Reader)
		skipped := 0
		for scanner.Scan() {
			skipped++
			if skipped >= s.NumOfLines {
				break
			}
		}

		r, w := io.Pipe()
		s.remain = r
		go func() {
			defer w.Close()
			for scanner.Scan() {
				w.Write(scanner.Bytes())
				w.Write([]byte("\n"))
			}
		}()
	})

	return s.remain.Read(p)
}
