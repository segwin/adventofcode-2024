package day15

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

type Tile byte

func (t Tile) String() string { return string(t) }

const (
	Empty = '.'
	Wall  = '#'
	Box   = 'O'
	BoxL  = '['
	BoxR  = ']'
	Robot = '@'
)

type Layout struct {
	map2d.Map[Tile]
}
