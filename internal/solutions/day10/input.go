package day10

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	for sc.Scan() {
		line := sc.Bytes()
		row := make([]int, len(line))
		for i, v := range line {
			height, err := strconv.Atoi(string(v))
			if err != nil {
				return nil, fmt.Errorf("parsing value %q as int: %w", v, err)
			}
			row[i] = height
		}
		s.Terrain = append(s.Terrain, row)
	}

	return &s, nil
}
