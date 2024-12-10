package solutions

import (
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
func RunAll() error {
	var errs []error
	for i, buildSolution := range Builders() {
		solution, err := buildSolution()
		if err != nil {
			return fmt.Errorf("building day %d solution: %w", i+1, err)
		}
		if err := solution.RunToConsole(); err != nil {
			errs = append(errs, fmt.Errorf("day %d: %w", i+1, err))
		}
	}
	return errors.Join(errs...)
}

// RunOne runs the given day's solution. The day param must be in [1, 25].
func RunOne(day int) error {
	solutions := Builders()
	if day < 1 || day > len(solutions) {
		return fmt.Errorf("%w (%d)", ErrInvalidDay, day)
	}

	solution, err := solutions[day-1]()
	if err != nil {
		return fmt.Errorf("building day %d solution: %w", day, err)
	}
	return solution.RunToConsole()
}

type Solution interface {
	// RunToConsole runs the solution and prints the result to the console.
	RunToConsole() error
}

func Builders() []func() (Solution, error) {
	return []func() (Solution, error){
		func() (Solution, error) { return day1.BuildSolution() },
		func() (Solution, error) { return day2.BuildSolution() },
		func() (Solution, error) { return day3.BuildSolution() },
		func() (Solution, error) { return day4.BuildSolution() },
		func() (Solution, error) { return day5.BuildSolution() },
		func() (Solution, error) { return day6.BuildSolution() },
		func() (Solution, error) { return day7.BuildSolution() },
		func() (Solution, error) { return day8.BuildSolution() },
		func() (Solution, error) { return day9.BuildSolution() },
		func() (Solution, error) { return day10.BuildSolution() },
		func() (Solution, error) { return day11.BuildSolution() },
		func() (Solution, error) { return day12.BuildSolution() },
		func() (Solution, error) { return day13.BuildSolution() },
		func() (Solution, error) { return day14.BuildSolution() },
		func() (Solution, error) { return day15.BuildSolution() },
		func() (Solution, error) { return day16.BuildSolution() },
		func() (Solution, error) { return day17.BuildSolution() },
		func() (Solution, error) { return day18.BuildSolution() },
		func() (Solution, error) { return day19.BuildSolution() },
		func() (Solution, error) { return day20.BuildSolution() },
		func() (Solution, error) { return day21.BuildSolution() },
		func() (Solution, error) { return day22.BuildSolution() },
		func() (Solution, error) { return day23.BuildSolution() },
		func() (Solution, error) { return day24.BuildSolution() },
		func() (Solution, error) { return day25.BuildSolution() },
	}
}
