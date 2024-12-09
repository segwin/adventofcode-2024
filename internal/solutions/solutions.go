package solutions

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrInvalidDay = errors.New("invalid day")
)

// RunAll runs each day's solution sequentially.
// If a day's solution fails, the next day is still executed.
func RunAll(ctx context.Context) error {
	var errs []error
	for i, solution := range all() {
		if err := solution.Run(ctx); err != nil {
			errs = append(errs, fmt.Errorf("day %d: %w", i+1, err))
		}
	}
	return errors.Join(errs...)
}

// RunOne runs the given day's solution. The day param must be in [1, 25].
func RunOne(ctx context.Context, day int) error {
	solutions := all()

	idx := day - 1
	if idx < 0 || idx > len(solutions) {
		return fmt.Errorf("%w (%d)", ErrInvalidDay, day)
	}

	return solutions[idx].Run(ctx)
}

type solution interface {
	Run(_ context.Context) error
}

func all() []solution {
	return []solution{
		&day1{},
		&day2{},
		&day3{},
		&day4{},
		&day5{},
		&day6{},
		&day7{},
		&day8{},
		&day9{},
		&day10{},
		&day11{},
		&day12{},
		&day13{},
		&day14{},
		&day15{},
		&day16{},
		&day17{},
		&day18{},
		&day19{},
		&day20{},
		&day21{},
		&day22{},
		&day23{},
		&day24{},
		&day25{},
	}
}
