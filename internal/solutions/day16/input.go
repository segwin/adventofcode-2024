package day16

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
		s.Maze = append(s.Maze, []Tile(sc.Text()))
	}

	return &s, nil
}
