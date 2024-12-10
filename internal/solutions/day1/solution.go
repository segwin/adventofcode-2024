package day1

import (
	"errors"
	"fmt"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

var (
	ErrMismatchedLens = errors.New("mismatched list lengths")
)

type Solution struct {
	// Left group's reported IDs.
	Left []int
	// Right group's reported IDs.
	Right []int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 1:\n")

	totalDistance, err := TotalDistance(s.Left, s.Right)
	if err != nil {
		return fmt.Errorf("computing total distance: %w", err)
	}
	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total distance: %d\n", totalDistance)

	similarity, err := Similarity(s.Left, s.Right)
	if err != nil {
		return fmt.Errorf("computing similarity score: %w", err)
	}
	fmt.Print("  PART 2:\n")
	fmt.Printf("    Similarity score: %d\n", similarity)

	return nil
}

func TotalDistance(left, right []int) (int, error) {
	if err := validateLens(left, right); err != nil {
		return 0, err
	}

	// sort lists to compare lowest values pairwise
	left = slices.Sorted(slices.Values(left))
	right = slices.Sorted(slices.Values(right))

	totalDistance := 0
	for i := range left {
		totalDistance += transform.Abs(left[i] - right[i])
	}

	return totalDistance, nil
}

func Similarity(left, right []int) (int, error) {
	if err := validateLens(left, right); err != nil {
		return 0, err
	}

	score := 0
	for _, lv := range left {
		for _, rv := range right {
			if lv == rv {
				score += lv
			}
		}
	}

	return score, nil
}

func validateLens(left, right []int) error {
	if len(left) != len(right) {
		return fmt.Errorf("%w (left = %d, right = %d)", ErrMismatchedLens, len(left), len(right))
	}
	return nil
}
