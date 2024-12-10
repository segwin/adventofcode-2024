package day1

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"

	"github.com/segwin/adventofcode-2024/internal/parsing"
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

		intRow, err := atois(row)
		if err != nil {
			return nil, fmt.Errorf("parsing columns as ints in input data: %w", err)
		}

		s.Left = append(s.Left, intRow[0])
		s.Right = append(s.Right, intRow[1])
	}

	return &s, nil
}

func atois(strs []string) ([]int, error) {
	ints := make([]int, len(strs))
	for i, str := range strs {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("column %d: %w", i, err)
		}
		ints[i] = v
	}
	return ints, nil
}
