package parsing

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"
)

var (
	ErrInvalidData = errors.New("invalid input data")
)

// SSVReader is a decoder for space-separated values.
type SSVReader struct {
	// Data to decode as space-separated values.
	Data io.Reader
	// ExpectedCols is the number of columns expected per row. If zero, the column check is skipped.
	ExpectedCols int
}

// All returns an iterator that reads one row at a time in the input data.
func (r *SSVReader) All() iter.Seq2[[]string, error] {
	return func(yield func([]string, error) bool) {
		s := bufio.NewScanner(r.Data)
		for s.Scan() {
			record := strings.Fields(s.Text())
			if r.ExpectedCols != 0 && len(record) != r.ExpectedCols {
				yield(nil, fmt.Errorf("%w: expected %d columns per row, got %d", ErrInvalidData, r.ExpectedCols, len(record)))
				return // invalid data, stop immediately
			}

			if !yield(record, nil) {
				return // caller has stopped iterating
			}
		}
	}
}
