package day3

import (
	"fmt"
	"regexp"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

type Solution struct {
	Memory string
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 3:\n")

	sum, err := SumMuls(s.Memory, false)
	if err != nil {
		return fmt.Errorf("executing muls in memory: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Sum: %d\n", sum)

	withConditionals, err := SumMuls(s.Memory, true)
	if err != nil {
		return fmt.Errorf("executing muls in memory with conditionals enabled: %w", err)
	}

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Sum with conditionals: %d\n", withConditionals)

	return nil
}

func SumMuls(memory string, withConditionals bool) (sum int, err error) {
	pattern := `(mul)\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`
	if withConditionals {
		pattern += `|do\(\)|don't\(\)`
	}
	rgx := regexp.MustCompile(pattern)

	mulsEnabled := true
	allMatches := rgx.FindAllStringSubmatch(memory, -1)
	for _, matches := range allMatches {
		switch matches[0] {
		case "do()":
			mulsEnabled = true
		case "don't()":
			mulsEnabled = false
		default: // mul
			if !mulsEnabled {
				continue
			}
			parsedVals, err := transform.Atois(matches[2], matches[3])
			if err != nil {
				// note: this shouldn't happen if the regex is well-formed, but let's be safe
				return 0, fmt.Errorf("parsing mul values (implementation error): %w", err)
			}
			sum += parsedVals[0] * parsedVals[1]
		}
	}
	return sum, nil
}
