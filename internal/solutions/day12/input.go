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
		line := sc.Text()
		s.Garden = append(s.Garden, []byte(line))
	}

	return &s, nil
}
