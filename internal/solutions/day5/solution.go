package day5

import (
	"fmt"
	"iter"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/parsing"
)

type Solution struct {
	// PagesAfter is the set of rules each update should respect.
	// Each entry's values is the set of pages that must come after the keyed page.
	PagesAfter map[int][]int

	// Updates is the set of pages to be published.
	Updates [][]int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 5:\n")

	// sanity check: updates must contain an odd number of pages to respect "middle page" req
	for i, pages := range s.Updates {
		if len(pages)%2 == 0 {
			return fmt.Errorf("%w: update %d has an even number of pages", parsing.ErrInvalidData, i)
		}
	}

	correctSum, err := SumCorrectUpdates(s.Updates, s.PagesAfter)
	if err != nil {
		return fmt.Errorf("summing correctly-ordered updates: %w", err)
	}
	fmt.Print("  PART 1:\n")
	fmt.Printf("    Middle page sum for correct: %d\n", correctSum)

	incorrectSum, err := SumIncorrectUpdates(s.Updates, s.PagesAfter)
	if err != nil {
		return fmt.Errorf("summing correctly-ordered updates: %w", err)
	}
	fmt.Print("  PART 2:\n")
	fmt.Printf("    Middle page sum for reordered incorrect updates: %d\n", incorrectSum)

	return nil
}

// SumCorrectUpdates returns the sum of each correctly-ordered update's middle page.
//
// Each update must contain an odd number of pages to avoid middle-page ambiguity.
func SumCorrectUpdates(updates [][]int, pagesAfter map[int][]int) (int, error) {
	sum := 0
	for update := range filterByCorrectness(updates, pagesAfter, true) {
		sum += middleElement(update)
	}
	return sum, nil
}

// SumIncorrectUpdates returns the sum of each incorrectly-ordered update's middle page after reordering.
//
// Each update must contain an odd number of pages to avoid middle-page ambiguity.
func SumIncorrectUpdates(updates [][]int, pagesAfter map[int][]int) (int, error) {
	isBefore := func(a, b int) int {
		if slices.Contains(pagesAfter[a], b) {
			return -1
		}
		if slices.Contains(pagesAfter[b], a) {
			return 1
		}
		return 0
	}

	sum := 0
	for update := range filterByCorrectness(updates, pagesAfter, false) {
		correctedUpdate := slices.SortedFunc(slices.Values(update), isBefore)
		sum += middleElement(correctedUpdate)
	}
	return sum, nil
}

// filterByCorrectness returns the subset of updates that match the desired correctness.
// Correctness is determined by checking the given set of page rules.
func filterByCorrectness(updates [][]int, pagesAfter map[int][]int, wantCorrect bool) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		for _, update := range updates {
			if isCorrect(update, pagesAfter) == wantCorrect {
				if !yield(update) {
					return
				}
			}
		}
	}
}

// isCorrect returns true if all pages in update respect the rules in pagesAfter.
func isCorrect(update []int, pagesAfter map[int][]int) bool {
	// check each page after 1st: preceding pages must not violate pagesAfter rules
	for i := 1; i < len(update); i++ {
		for _, pageAfter := range pagesAfter[update[i]] {
			if slices.Contains(update[:i], pageAfter) {
				return false // ordering not respected
			}
		}
	}
	return true
}

func middleElement(s []int) int {
	return s[len(s)/2]
}
