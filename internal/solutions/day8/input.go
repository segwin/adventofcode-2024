package day8

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
		row := make([]Tile, len(line))
		for i, t := range line {
			row[i] = Tile(t)
		}
		s.CityMap = append(s.CityMap, row)
	}

	return &s, nil
}
