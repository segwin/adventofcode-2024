package map2d

import "slices"

// Map of 2-dimensional values.
type Map[T comparable] [][]T

// Get the value at the given position. If that position is outside the map, ok is set to false.
func (m Map[T]) Get(pos Position) (value T, ok bool) {
	if pos.Y < 0 || pos.Y >= len(m) || pos.X < 0 || pos.X >= len(m[pos.Y]) {
		var zero T
		return zero, false
	}
	return m[pos.Y][pos.X], true
}

// Find returns the first position with the given value, if any.
func (m Map[T]) Find(v T) (pos Position, ok bool) {
	for i, row := range m {
		for j, cell := range row {
			if cell == v {
				return PositionFromIndex(i, j), true
			}
		}
	}
	return Position{}, false
}

// FindAll returns all positions with the given value, if any.
func (m Map[T]) FindAll(v T) []Position {
	var found []Position
	for i, row := range m {
		for j, cell := range row {
			if cell == v {
				found = append(found, PositionFromIndex(i, j))
			}
		}
	}
	return found
}

// Contains returns true if the given position is present in the map.
//
// Equivalent to _, contains := m.Get(pos).
func (m Map[T]) Contains(pos Position) bool {
	_, ok := m.Get(pos)
	return ok
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
