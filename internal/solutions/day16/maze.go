package day16

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

// Tile in the maze.
type Tile byte

const (
	Wall  = '#'
	Empty = '.'
	Start = 'S'
	End   = 'E'
)

// Maze as given by the puzzle map.
type Maze = map2d.Map[Tile]
