package day14

import "fmt"

type Solution struct {
	Robots []RobotState
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 14:\n")

	layout := NewMap(101, 103) // dimensions from problem statement
	after100s := After(100, s.Robots, layout)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Safety factor: %d\n", SafetyFactor(after100s, layout))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// After returns the new set of robot states after the given number of seconds has elapsed.
func After(seconds int, initialStates []RobotState, layout Map) []RobotState {
	newStates := make([]RobotState, len(initialStates))
	for i, initialState := range initialStates {
		newStates[i] = initialState.After(seconds, layout)
	}
	return newStates
}

// SafetyFactor returns the product of robot counts in each quadrant.
func SafetyFactor(robots []RobotState, layout Map) int {
	product := 1
	for _, count := range countByQuadrant(robots, layout) {
		product *= count
	}
	return product
}

// countByQuadrant returns the number of robots present in each quadrant.
func countByQuadrant(robots []RobotState, layout Map) (counts [4]int) {
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
