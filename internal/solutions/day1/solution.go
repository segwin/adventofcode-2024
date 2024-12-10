package day1

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrMismatchedLens = errors.New("mismatched list lengths")
)

type Solution struct {
	Left  []int
	Right []int
}

func (s *Solution) RunToConsole() error {
	totalDistance, err := s.TotalDistance()
	if err != nil {
		return fmt.Errorf("computing total distance: %w", err)
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
	return nil
}

func (s *Solution) TotalDistance() (int, error) {
	if len(s.Left) != len(s.Right) {
		return 0, fmt.Errorf("%w (left = %d, right = %d)", ErrMismatchedLens, len(s.Left), len(s.Right))
	}

	slices.Sort(s.Left)
	slices.Sort(s.Right)
	totalDistance := 0
	for i := range s.Left {
		low, high := s.Left[i], s.Right[i]
		if high < low {
			low, high = high, low
		}
		totalDistance += high - low
	}

	return totalDistance, nil
}
