package day18

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

type Layout = map2d.Map[Tile]

// NewEmptyLayout creates a new memory layout with the given dimensions. All tiles are set to Empty.
func NewEmptyLayout(lenX, lenY int) Layout {
	m := map2d.NewMap[Tile](lenX, lenY)
	for i, row := range m {
		for j := range row {
			m[i][j] = Empty
		}
	}
	return m
}

type Tile byte

const (
	Empty     Tile = '.'
	Corrupted Tile = '#'
)

func (t Tile) String() string { return string(t) }
