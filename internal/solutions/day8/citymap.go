package day8

type Tile byte

const (
	Empty Tile = '.'
)

type CityMap [][]Tile

func (m CityMap) Contains(p Position) bool {
	return p.Y >= 0 && p.Y < len(m) && p.X >= 0 && p.X < len(m[p.Y])
}

type Position struct {
	X, Y int
}

func (p Position) Sub(p2 Position) Position {
	return Position{X: p.X - p2.X, Y: p.Y - p2.Y}
}

func (p Position) Add(p2 Position) Position {
	return Position{X: p.X + p2.X, Y: p.Y + p2.Y}
}
