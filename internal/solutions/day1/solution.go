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

	totalDistance, err := s.TotalDistance()
	if err != nil {
		return fmt.Errorf("computing total distance: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total distance: %d\n", totalDistance)

	similarity, err := s.Similarity()
	if err != nil {
		return fmt.Errorf("computing similarity score: %w", err)
	}

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Similarity score: %d\n", similarity)
	return nil
}

func (s *Solution) TotalDistance() (int, error) {
	if err := s.validateInputs(); err != nil {
		return 0, err
	}

	// sort lists to compare lowest values pairwise
	left := slices.Sorted(slices.Values(s.Left))
	right := slices.Sorted(slices.Values(s.Right))

	totalDistance := 0
	for i := range left {
		totalDistance += transform.Abs(left[i] - right[i])
	}

	return totalDistance, nil
}

func (s *Solution) Similarity() (int, error) {
	if err := s.validateInputs(); err != nil {
		return 0, err
	}

	score := 0
	for _, lv := range s.Left {
		for _, rv := range s.Right {
			if lv == rv {
				score += lv
			}
		}
	}

	return score, nil
}

func (s *Solution) validateInputs() error {
	if len(s.Left) != len(s.Right) {
		return fmt.Errorf("%w (left = %d, right = %d)", ErrMismatchedLens, len(s.Left), len(s.Right))
	}
	return nil
}
