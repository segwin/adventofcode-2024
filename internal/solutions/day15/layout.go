package day15

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

type Tile byte

const (
	Empty = '.'
	Wall  = '#'
	Box   = 'O'
	Robot = '@'
)

type Layout struct {
	map2d.Map[Tile]
}
