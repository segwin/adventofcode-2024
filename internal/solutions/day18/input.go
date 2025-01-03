package day18

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"errors"
	"fmt"
	"io"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	r := csv.NewReader(bytes.NewReader(inputData))
	for {
		records, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("%w: %w", parsing.ErrInvalidData, err)
		}

		if len(records) != 2 {
			return nil, fmt.Errorf("%w: expected 2 entries per row, got %d", parsing.ErrInvalidData, len(records))
		}

		xy, err := transform.Atois(records...)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", parsing.ErrInvalidData, err)
		}

		s.FallingBytes = append(s.FallingBytes, map2d.Position{X: xy[0], Y: xy[1]})
	}

	return &s, nil
}
