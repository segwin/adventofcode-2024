package map2d

// Direction for a movement on the map grid.
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
