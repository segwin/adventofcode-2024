package day15

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Layout       Layout
	Instructions []map2d.Direction
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 15:\n")

	finalLayout, _ := Resolve(s.Layout, s.Instructions...)
	boxes := finalLayout.FindAll(Box)

	gpsSum := 0
	for _, box := range boxes {
		gpsSum += 100*box.Y + box.X
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Box GPS sum: %d\n", gpsSum)

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// Resolve all instructions for the robot, returning the final updated layout.
// Returns an error if the robot can't be found in layout.
func Resolve(layout Layout, instructions ...map2d.Direction) (Layout, error) {
	robot, ok := layout.Find(Robot)
	if !ok {
		return Layout{}, fmt.Errorf("%w: robot not found", parsing.ErrInvalidData)
	}

	for _, instruction := range instructions {
		layout, robot = resolveOne(layout, robot, instruction)
	}
	return layout, nil
}

func resolveOne(curLayout Layout, curRobot map2d.Position, instruction map2d.Direction) (newLayout Layout, newRobot map2d.Position) {
	newRobot = curRobot.Move(instruction, 1)

	nextPos := newRobot
	nextTile, _ := curLayout.Get(nextPos) // next tile should never be out of bounds

	var nextBoxes []map2d.Position
	for nextTile == Box {
		nextBoxes = append(nextBoxes, nextPos)
		nextPos = nextPos.Move(instruction, 1)
		nextTile, _ = curLayout.Get(nextPos) // next tile should never be out of bounds
	}

	if nextTile == Wall {
		return curLayout, curRobot // do nothing: can't move walls
	}

	// found an empty tile: move robot ahead & move any boxes that were in the way
	curLayout.Map = curLayout.With(curRobot, Empty)
	curLayout.Map = curLayout.With(newRobot, Robot)
	for _, curBox := range nextBoxes {
		curLayout.Map = curLayout.With(curBox.Move(instruction, 1), Box)
	}

	return curLayout, newRobot
}
