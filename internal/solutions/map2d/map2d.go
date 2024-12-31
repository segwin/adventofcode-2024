package map2d

import "slices"

// Map of 2-dimensional values.
type Map[T any] [][]T

// Get the value at the given position. If that position is outside the map, ok is set to false.
func (m Map[T]) Get(pos Position) (value T, ok bool) {
	if pos.Y < 0 || pos.Y >= len(m) || pos.X < 0 || pos.X >= len(m[pos.Y]) {
		var zero T
		return zero, false
	}
	return m[pos.Y][pos.X], true
}

// With returns a new map with the given position set to newValue.
// Unrelated rows are shared in memory with the original map.
//
// No bounds checking is performed on pos.
func (m Map[T]) With(pos Position, newValue T) Map[T] {
	// create copy of map with a new row allocated before mutating it
	newMap := slices.Clone(m)
	newMap[pos.Y] = slices.Clone(newMap[pos.Y])

	// update the value at the given position
	newMap[pos.Y][pos.X] = newValue
	return newMap
}

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

// LessThan returns true if this position is less than p2. Ordering is implementation-defined.
func (p Position) LessThan(p2 Position) bool {
	if p.X < p2.X {
		return true
	} else if p.X > p2.X {
		return false
	}
	if p.Y < p2.Y {
		return true
	} else if p.Y > p2.Y {
		return false
	}
	return false // equal
}

// PositionFromIndex creates a Position object from the given map indices.
// Useful since Position represents visual X/Y coordinates which are the inverse of i/j indices.
func PositionFromIndex(i, j int) Position {
	return Position{X: j, Y: i}
}
