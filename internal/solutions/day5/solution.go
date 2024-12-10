package day5

import (
	"fmt"
	"iter"
	"slices"

	"github.com/segwin/adventofcode-2024/internal/parsing"
)

type Solution struct {
	// Rules for page ordering each update should respect.
	Rules []PageRule
	// Updates is the set of pages to be published.
	Updates [][]int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 5:\n")

	correctSum, err := SumCorrectUpdates(s.Updates, s.Rules)
	if err != nil {
		return fmt.Errorf("summing correctly-ordered updates: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Middle page sum: %d\n", correctSum)
	return nil
}

// SumCorrectUpdates returns the sum of each correctly-ordered update's middle page.
//
// Returns an error if any update contains an even number of pages (ambiguous middle page).
func SumCorrectUpdates(updates [][]int, rules []PageRule) (int, error) {
	// sanity check: updates must contain an odd number of pages to respect "middle page" req
	for i, pages := range updates {
		if len(pages)%2 == 0 {
			return 0, fmt.Errorf("%w: update %d has an even number of pages", parsing.ErrInvalidData, i)
		}
	}

	sum := 0
	for midPage := range middlePages(correctlyOrderedUpdates(updates, rules)) {
		sum += midPage
	}
	return sum, nil
}

func correctlyOrderedUpdates(updates [][]int, rules []PageRule) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		pagesAfter := AggregateRules(rules)
		for _, update := range updates {
			if isCorrectlyOrdered(update, pagesAfter) {
				if !yield(update) {
					return
				}
			}
		}
	}
}

func isCorrectlyOrdered(update []int, pagesAfter map[int][]int) bool {
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

func middlePages(updates iter.Seq[[]int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for pages := range updates {
			if !yield(pages[len(pages)/2]) {
				return
			}
		}
	}
}
