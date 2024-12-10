package day3

import (
	_ "embed"
)

//go:embed input.txt
var inputData string

func BuildSolution() (*Solution, error) {
	return &Solution{Memory: inputData}, nil
}
