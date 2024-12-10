package day2

import (
	"fmt"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

type Report []int

type Solution struct {
	// Reports from the engineers. Each value is a level.
	Reports []Report
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 2:\n")

	safeReports, err := SafeReports(s.Reports, false)
	if err != nil {
		return fmt.Errorf("counting safe reports: %w", err)
	}
	fmt.Print("  PART 1:\n")
	fmt.Printf("    Safe reports: %d\n", safeReports)

	withDampening, err := SafeReports(s.Reports, true)
	if err != nil {
		return fmt.Errorf("counting safe reports with problem dampening: %w", err)
	}
	fmt.Print("  PART 2:\n")
	fmt.Printf("    Safe reports after problem dampening: %d\n", withDampening)

	return nil
}

// SafeReports counts the number of "safe" engineering reports.
//
// A report is safe if all of the following conditions are met:
//   - the report is monotonic (strictly increasing or decreasing)
//   - adjacent report values only differ by 1 or 2
//
// If problemDampening is true, up to one "unsafe" level will be ignored.
func SafeReports(reports []Report, problemDampening bool) (int, error) {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++ // safe
			continue
		}

		if !problemDampening {
			continue // unsafe & no problem dampening
		}

		for i := 0; i < len(report); i++ {
			if isSafe(withoutElement(report, i)) {
				count++ // safe after problem dampening
				break
			}
		}
	}

	return count, nil
}

func isSafe(report Report) bool {
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

func withoutElement[T any](slice []T, i int) []T {
	return append(slices.Clone(slice[:i]), slice[i+1:]...)
}

type direction int

const (
	decreasing direction = -1
	increasing direction = +1
)
