package day12

import (
	"bufio"
	"bytes"
	_ "embed"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	for sc.Scan() {
		line := sc.Bytes()
		row := make([]byte, len(line))
		copy(row, line)
		s.Garden = append(s.Garden, row)
	}

	return &s, nil
}
