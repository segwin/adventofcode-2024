package day6

import (
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

// GuardState wandering around the map using a strict ruleset.
type GuardState struct {
	// Position of the guard on the map grid.
	Position map2d.Position
	// Direction the guard is currently walking in. Changes when encountering an obstruction.
	Direction Direction
}

// AdvanceOne makes the guard take a step forward or turn 90 degrees based on their current state,
// returning their new state afterwards.
//
// If the guard exits the room, the returned state is nil.
func (g GuardState) AdvanceOne(floorMap FloorMap) (*GuardState, FloorMap) {
	// advance guard to next position
	nextPosition := map2d.Position{
		X: g.Position.X + g.Direction.x,
		Y: g.Position.Y + g.Direction.y,
	}
	nextTile, ok := floorMap.Get(nextPosition)
	if !ok {
		// the guard has left the room
		return nil, floorMap
	}

	if nextTile == Obstacle {
		// something's blocking the path, turn right 90 degrees
		newState := &GuardState{Position: g.Position, Direction: g.Direction.turnRight()}
		return newState, floorMap.With(g.Position, newState.Tile())
	}

	// nothing ahead, take a step
	newState := &GuardState{Position: nextPosition, Direction: g.Direction}
	return newState, floorMap.With(newState.Position, newState.Tile())
}

func (g GuardState) Tile() Tile {
	switch g.Direction {
	case North():
		return GuardNorth
	case East():
		return GuardEast
	case South():
		return GuardSouth
	case West():
		return GuardWest
	}
	return Empty
}

// Direction for a movement on the map grid.
type Direction struct {
	x int
	y int
}

func North() Direction { return Direction{y: -1} }
func East() Direction  { return Direction{x: 1} }
func South() Direction { return Direction{y: 1} }
func West() Direction  { return Direction{x: -1} }

// TurnRight returns a new direction after a 90 degree rotation clockwise.
func (d Direction) turnRight() Direction {
	// we could use trigonometry instead, but since we're at right angles we can easily switch-case
	// it and avoid rounding errors & angle normalization
	switch d {
	case East():
		return South()
	case South():
		return West()
	case West():
		return North()
	case North():
		return East()
	default:
		panic("implementation error: non-cardinal Direction")
	}
}

// Tile on the map grid.
type Tile byte

const (
	Empty      Tile = '.'
	Obstacle   Tile = '#'
	GuardEast  Tile = '>'
	GuardNorth Tile = '^'
	GuardWest  Tile = '<'
	GuardSouth Tile = 'v'
)

// IsGuard returns true if this tile represents a guard's state, past or present.
func (t Tile) IsGuard() bool {
	switch t {
	case GuardNorth, GuardEast, GuardSouth, GuardWest:
		return true
	default:
		return false
	}
}

// Direction the guard is or was moving on this tile. If this tile isn't a guard's state, returns
// an empty direction instead.
func (t Tile) Direction() Direction {
	switch t {
	case GuardNorth:
		return North()
	case GuardEast:
		return East()
	case GuardSouth:
		return South()
	case GuardWest:
		return West()
	}
	return Direction{} // others: no direction
}

// FloorMap is the map of the room.
type FloorMap = map2d.Map[Tile]
