package day13

import "github.com/segwin/adventofcode-2024/internal/solutions/map2d"

const (
	// ACost is the cost to push the A button, in tokens.
	ACost = 3
	// BCost is the cost to push the B button, in tokens.
	BCost = 1
)

type ClawMachine struct {
	// MoveA is the movement effect from pressing button MoveA.
	MoveA map2d.Distance
	// MoveB is the movement effect from pressing button MoveB.
	MoveB map2d.Distance
	// Prize position within the claw machine, measured as distance from the origin.
	Prize map2d.Position
}
