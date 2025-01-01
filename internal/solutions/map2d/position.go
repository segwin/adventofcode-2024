package map2d

// Position on a 2D Map.
type Position struct {
	X int
	Y int
}

// Sub returns the difference between this position and p2.
func (p Position) Sub(p2 Position) Position {
	return Position{X: p.X - p2.X, Y: p.Y - p2.Y}
}

// Add returns the sum of this position and p2.
func (p Position) Add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}

// Move in the given direction by the given amount.
func (p Position) Move(d Direction, amount int) Position {
	return p.Add(Position{X: amount * d.offset.X, Y: amount * d.offset.Y})
}

// LessThan returns true if this position is less than p2. Ordering is implementation-defined.
func (p Position) LessThan(p2 Position) bool {
	if p.X < p2.X {
		return true
	}
	if p.X > p2.X {
		return false
	}
	if p.Y < p2.Y {
		return true
	}
	return false // equal or greater
}

// ProjectOnto returns the component of p that aligns with the given direction.
func (p Position) ProjectOnto(d Direction) int {
	switch d {
	case North():
		return -p.Y
	case East():
		return p.X
	case South():
		return p.Y
	case West():
		return -p.X
	default:
		return 0 // invalid or zero direction: x*0 == 0
	}
}

// PositionFromIndex creates a Position object from the given map indices.
// Useful since Position represents visual X/Y coordinates which are the inverse of i/j indices.
func PositionFromIndex(i, j int) Position {
	return Position{X: j, Y: i}
}
