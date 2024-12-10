package day4

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var inputData string

func BuildSolution() (*Solution, error) {
	return &Solution{
		Search: strings.Split(inputData, "\n"),
	}, nil
}
