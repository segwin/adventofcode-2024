package day2

import (
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

type Solution struct {
	// Reports from the engineers. Each value is a level.
	Reports [][]int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 2:\n")

	safeReports, err := s.SafeReports()
	if err != nil {
		return fmt.Errorf("counting safe reports: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Safe reports: %d\n", safeReports)
	return nil
}

func (s *Solution) SafeReports() (int, error) {
	count := 0
	for _, report := range s.Reports {
		if isSafe(report) {
			count++
		}
	}
	return count, nil
}

// isSafe returns false in any of the following cases:
//   - the report is not monotonic (strictly increasing or decreasing)
//   - 2+ adjacent report values differ by <1 or >3
func isSafe(report []int) bool {
	if len(report) < 2 {
		return true // all conditions are met if report does not contain multiple values
	}

	var lastDirection direction
	for i := 0; i < len(report)-1; i++ {
		cur := report[i]
		next := report[i+1]

		// check adjacent value distance
		delta := next - cur
		if absDelta := transform.Abs(float64(delta)); absDelta < 1 || absDelta > 3 {
			return false
		}

		// check monotonicity
		var newDirection direction
		if delta < 0 {
			newDirection = decreasing
		} else {
			newDirection = increasing
		}
		if lastDirection == 0 { // not determined yet: set it now
			lastDirection = newDirection
			continue
		}
		if newDirection != lastDirection {
			return false
		}
	}

	return true
}

type direction int

const (
	decreasing direction = -1
	increasing direction = +1
)
