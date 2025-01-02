package day14

import (
	"bytes"
	"fmt"
	"os"

	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type Solution struct {
	Robots []RobotState
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 14:\n")

	layout := NewLayout(101, 103) // dimensions from problem statement
	after100s := After(100, s.Robots, layout)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Safety factor: %d\n", SafetyFactor(after100s, layout))

	outputFile, err := os.CreateTemp(os.TempDir(), "day14_part2")
	if err != nil {
		return fmt.Errorf("creating temporary output file: %w", err)
	}

	fmt.Printf("  PART 2: rendering 10000 iterations to %s\n", outputFile.Name())
	for i := 1; i <= 10000; i++ {
		rows := RenderMap(After(i, s.Robots, layout), layout)
		fmt.Fprintf(outputFile, "Iteration %d:\n%s\n\n", i, bytes.Join(rows, []byte{'\n'}))
	}

	return nil
}

// After returns the new set of robot states after the given number of seconds has elapsed.
func After(seconds int, initialStates []RobotState, layout Layout) []RobotState {
	newStates := make([]RobotState, len(initialStates))
	for i, initialState := range initialStates {
		newStates[i] = initialState.After(seconds, layout)
	}
	return newStates
}

// SafetyFactor returns the product of robot counts in each quadrant.
func SafetyFactor(robots []RobotState, layout Layout) int {
	product := 1
	for _, count := range countByQuadrant(robots, layout) {
		product *= count
	}
	return product
}

func RenderMap(robots []RobotState, layout Layout) [][]byte {
	mapArea := len(layout) * len(layout[0])
	robotPositions := make(map[map2d.Position]struct{}, mapArea) // worst case: 1 robot per tile
	for _, robot := range robots {
		robotPositions[robot.Position] = struct{}{}
	}

	msgRows := make([][]byte, len(layout))
	for i, row := range layout {
		msgRows[i] = make([]byte, len(row))
		for j := range row {
			msgRows[i][j] = ' '
			if _, ok := robotPositions[map2d.PositionFromIndex(i, j)]; ok {
				msgRows[i][j] = '*'
			}
		}
	}

	return msgRows
}

// countByQuadrant returns the number of robots present in each quadrant.
func countByQuadrant(robots []RobotState, layout Layout) (counts [4]int) {
	quadrants := layout.Quadrants()
	for _, robot := range robots {
		for i, quadrant := range quadrants {
			if quadrant.Contains(robot.Position) {
				counts[i]++
				break // robot can't be in 2 places at once
			}
		}
	}
	return counts
}
