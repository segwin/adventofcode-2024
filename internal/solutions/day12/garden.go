package day12

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

type Garden = map2d.Map[byte]

type Region struct {
	Kind      byte
	Area      int
	Perimeter int
}
