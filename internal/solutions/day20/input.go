package day20

import (
	"bytes"
	_ "embed"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	r := bytes.NewReader(inputData)
	return &Solution{
		Racetrack: map2d.DecodeMap(r, func(cell byte) Tile { return Tile(cell) }),
	}, nil
}
