package day19

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/parsing"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))

	// parse towels
	if !sc.Scan() {
		return nil, fmt.Errorf("%w: missing towels line", parsing.ErrInvalidData)
	}
	s.Towels = strings.Split(sc.Text(), ", ")

	// parse designs
	if !sc.Scan() {
		return nil, fmt.Errorf("%w: missing newline betweeen towels & designs", parsing.ErrInvalidData)
	}
	for sc.Scan() {
		s.Designs = append(s.Designs, sc.Text())
	}

	return &s, nil
}
