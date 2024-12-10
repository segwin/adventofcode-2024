package day1

import (
	"bytes"
	_ "embed"
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	r := parsing.SSVReader{Data: bytes.NewReader(inputData), ExpectedCols: 2}
	for row, err := range r.All() {
		if err != nil {
			return nil, fmt.Errorf("parsing input data: %w", err)
		}

		intRow, err := transform.Atois(row)
		if err != nil {
			return nil, err
		}

		s.Left = append(s.Left, intRow[0])
		s.Right = append(s.Right, intRow[1])
	}

	return &s, nil
}
