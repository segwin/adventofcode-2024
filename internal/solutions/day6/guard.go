package day6

import (
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

// GuardState wandering around the map using a strict ruleset.
type GuardState struct {
	// Position of the guard on the map grid.
	Position map2d.Position
	// Direction the guard is currently walking in. Changes when encountering an obstruction.
	Direction map2d.Direction
}

// AdvanceOne makes the guard take a step forward or turn 90 degrees based on their current state,
// returning their new state afterwards.
//
// If the guard exits the room, the returned state is nil.
func (g GuardState) AdvanceOne(floorMap FloorMap) (*GuardState, FloorMap) {
	// advance guard to next position
	nextPosition := g.Position.Move(g.Direction, 1)
	nextTile, ok := floorMap.Get(nextPosition)
	if !ok {
		// the guard has left the room
		return nil, floorMap
	}

	if nextTile == Obstacle {
		// something's blocking the path, turn right 90 degrees
		newState := &GuardState{Position: g.Position, Direction: g.Direction.TurnClockwise()}
		return newState, floorMap.With(g.Position, newState.Tile())
	}

	// nothing ahead, take a step
	newState := &GuardState{Position: nextPosition, Direction: g.Direction}
	return newState, floorMap.With(newState.Position, newState.Tile())
}

func (g GuardState) Tile() Tile {
	switch g.Direction {
	case map2d.North():
		return GuardNorth
	case map2d.East():
		return GuardEast
	case map2d.South():
		return GuardSouth
	case map2d.West():
		return GuardWest
	}
	return Empty
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
func (t Tile) Direction() map2d.Direction {
	switch t {
	case GuardNorth:
		return map2d.North()
	case GuardEast:
		return map2d.East()
	case GuardSouth:
		return map2d.South()
	case GuardWest:
		return map2d.West()
	}
	return map2d.Direction{} // others: no direction
}

// FloorMap is the map of the room.
type FloorMap = map2d.Map[Tile]
