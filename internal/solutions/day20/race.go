package day20

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

// Tile on the racetrack.
type Tile byte

const (
	Wall  Tile = '#'
	Track Tile = '.'
	Start Tile = 'S'
	End   Tile = 'E'
)

// Cheat used during the race to glitch through a wall.
type Cheat struct {
	// Start and End positions for the cheat. End must always be an absolute distance of 1 from Start.
	Start, End map2d.Position
}
