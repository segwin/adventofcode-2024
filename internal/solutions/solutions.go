package solutions

import (
	"context"
	"errors"
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/solutions/day1"
	"github.com/segwin/adventofcode-2024/internal/solutions/day10"
	"github.com/segwin/adventofcode-2024/internal/solutions/day11"
	"github.com/segwin/adventofcode-2024/internal/solutions/day12"
	"github.com/segwin/adventofcode-2024/internal/solutions/day13"
	"github.com/segwin/adventofcode-2024/internal/solutions/day14"
	"github.com/segwin/adventofcode-2024/internal/solutions/day15"
	"github.com/segwin/adventofcode-2024/internal/solutions/day16"
	"github.com/segwin/adventofcode-2024/internal/solutions/day17"
	"github.com/segwin/adventofcode-2024/internal/solutions/day18"
	"github.com/segwin/adventofcode-2024/internal/solutions/day19"
	"github.com/segwin/adventofcode-2024/internal/solutions/day2"
	"github.com/segwin/adventofcode-2024/internal/solutions/day20"
	"github.com/segwin/adventofcode-2024/internal/solutions/day21"
	"github.com/segwin/adventofcode-2024/internal/solutions/day22"
	"github.com/segwin/adventofcode-2024/internal/solutions/day23"
	"github.com/segwin/adventofcode-2024/internal/solutions/day24"
	"github.com/segwin/adventofcode-2024/internal/solutions/day25"
	"github.com/segwin/adventofcode-2024/internal/solutions/day3"
	"github.com/segwin/adventofcode-2024/internal/solutions/day4"
	"github.com/segwin/adventofcode-2024/internal/solutions/day5"
	"github.com/segwin/adventofcode-2024/internal/solutions/day6"
	"github.com/segwin/adventofcode-2024/internal/solutions/day7"
	"github.com/segwin/adventofcode-2024/internal/solutions/day8"
	"github.com/segwin/adventofcode-2024/internal/solutions/day9"
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
	if day < 1 || day > len(solutions) {
		return fmt.Errorf("%w (%d)", ErrInvalidDay, day)
	}
	return solutions[day-1].Run(ctx)
}

type Solution interface {
	Run(context.Context) error
}

func all() []Solution {
	return []Solution{
		&day1.Solution{},
		&day2.Solution{},
		&day3.Solution{},
		&day4.Solution{},
		&day5.Solution{},
		&day6.Solution{},
		&day7.Solution{},
		&day8.Solution{},
		&day9.Solution{},
		&day10.Solution{},
		&day11.Solution{},
		&day12.Solution{},
		&day13.Solution{},
		&day14.Solution{},
		&day15.Solution{},
		&day16.Solution{},
		&day17.Solution{},
		&day18.Solution{},
		&day19.Solution{},
		&day20.Solution{},
		&day21.Solution{},
		&day22.Solution{},
		&day23.Solution{},
		&day24.Solution{},
		&day25.Solution{},
	}
}
