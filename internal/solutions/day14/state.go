package day14

import (
	"math"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Layout map2d.Map[struct{}]

// NewLayout creates a new Layout with the given X and Y dimensions.
func NewLayout(x, y int) Layout {
	m := make(Layout, y)
	for i := range m {
		m[i] = make([]struct{}, x)
	}
	return m
}

// Quadrants of the map. If the map has odd-numbered dimensions, the odd middle line is omitted.
func (m Layout) Quadrants() [4]Quadrant {
	maxX, maxY := len(m[0]), len(m)
	midX, midY := float64(maxX)/2, float64(maxY)/2

	return [4]Quadrant{
		{ // top-left
			Start: map2d.Position{X: 0, Y: 0},
			End:   map2d.Position{X: int(midX), Y: int(midY)},
		},
		{ // top-right
			Start: map2d.Position{X: int(math.Ceil(midX)), Y: 0},
			End:   map2d.Position{X: maxX, Y: int(midY)},
		},
		{ // bottom-left
			Start: map2d.Position{X: 0, Y: int(math.Ceil(midY))},
			End:   map2d.Position{X: int(midX), Y: maxY},
		},
		{ // bottom-right
			Start: map2d.Position{X: int(math.Ceil(midX)), Y: int(math.Ceil(midY))},
			End:   map2d.Position{X: maxX, Y: maxY},
		},
	}
}

// Quadrant of a Layout.
type Quadrant struct {
	Start map2d.Position
	End   map2d.Position
}

// Contains returns true if the given position exists in this quadrant.
func (q Quadrant) Contains(pos map2d.Position) bool {
	return q.Start.X <= pos.X && q.End.X > pos.X && q.Start.Y <= pos.Y && q.End.Y > pos.Y
}

// Velocity of an actor on a 2D Map. See Position for coordinate definitions.
type Velocity struct {
	X, Y int
}

// RobotState is the state of a robot on the map at a given point in time.
type RobotState struct {
	// Position of the robot on the map, in tiles.
	Position map2d.Position
	// Velocity of the robot, in tiles/second.
	Velocity Velocity
}

// After produces a new RobotState after the given number of seconds, with its position updated based
// on its velocity.
func (s RobotState) After(seconds int, layout Layout) RobotState {
	newState := RobotState{
		Position: s.Position.Add(s.Velocity.X*seconds, s.Velocity.Y*seconds),
		Velocity: s.Velocity,
	}

	// if robot exceeds maxX or maxY, it teleports back within bounds
	if maxX := len(layout[0]); newState.Position.X < 0 || newState.Position.X > maxX {
		newState.Position.X = wrapToPositive(newState.Position.X, maxX)
	}
	if maxY := len(layout); newState.Position.Y < 0 || newState.Position.Y > maxY {
		newState.Position.Y = wrapToPositive(newState.Position.Y, maxY)
	}

	return newState
}

func wrapToPositive(v, upper int) int {
	remainder := v % upper
	if remainder >= 0 {
		return remainder
	}
	return remainder + upper
}
