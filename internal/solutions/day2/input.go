package day2

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

	r := parsing.SSVReader{Data: bytes.NewReader(inputData)}
	for row, err := range r.All() {
		if err != nil {
			return nil, fmt.Errorf("parsing input data: %w", err)
		}

		intRow, err := transform.Atois(row)
		if err != nil {
			return nil, fmt.Errorf("parsing columns as ints in input data: %w", err)
		}

		s.Reports = append(s.Reports, intRow)
	}

	return &s, nil
}
