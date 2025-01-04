package map2d

// Direction for a movement along one of the map grid's 2 axes.
type Direction struct {
	offset Position
}

// North points upward on the map.
func North() Direction { return Direction{offset: Position{Y: -1}} }

// East points right on the map.
func East() Direction { return Direction{offset: Position{X: 1}} }

// South points downward on the map.
func South() Direction { return Direction{offset: Position{Y: 1}} }

// West points left on the map.
func West() Direction { return Direction{offset: Position{X: -1}} }

// CardinalDirections returns the full set of compass directions.
func CardinalDirections() []Direction {
	return []Direction{North(), East(), South(), West()}
}

// TurnClockwise returns a new direction after a 90 degree rotation clockwise.
func (d Direction) TurnClockwise() Direction {
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
		return d // invalid or zero direction: rotations have no effect
	}
}

// IsVertical returns true if this direction is North or South.
func (d Direction) IsVertical() bool { return d.offset.Y != 0 }

// IsHorizontal returns true if this direction is East or West.
func (d Direction) IsHorizontal() bool { return d.offset.X != 0 }

// Angle returns the angle between these two directions, in degrees.
// Returned values are one of: 0, 90, 180
func (d Direction) Angle(other Direction) int {
	switch {
	case d == other:
		return 0 // same direction
	case d.IsVertical() && other.IsVertical(), d.IsHorizontal() && other.IsHorizontal():
		return 180 // same axis but not the same => opposite directions
	default:
		return 90 // other combinations: 90 degrees (only cardinal directions are supported)
	}
}
