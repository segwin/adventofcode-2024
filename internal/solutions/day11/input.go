package day11

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	values, err := transform.Atois(strings.Fields(string(inputData))...)
	if err != nil {
		return nil, fmt.Errorf("parsing input data: %w", err)
	}

	s := Solution{InitialStones: make([]Stone, len(values))}
	for i, v := range values {
		s.InitialStones[i] = Stone(v)
	}

	return &s, nil
}
