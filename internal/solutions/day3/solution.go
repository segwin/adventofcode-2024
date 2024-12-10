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

	sum, err := SumMuls(s.Memory)
	if err != nil {
		return fmt.Errorf("executing muls in memory: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Sum: %d\n", sum)
	return nil
}

func SumMuls(memory string) (sum int, err error) {
	pattern := regexp.MustCompile(`mul\(([1-9][0-9]{0,2}),([1-9][0-9]{0,2})\)`)
	for _, matches := range pattern.FindAllStringSubmatch(memory, -1) {
		parsedVals, err := transform.Atois(matches[1], matches[2])
		if err != nil {
			// note: this shouldn't happen if the regex is well-formed, but let's be safe
			return 0, fmt.Errorf("parsing mul values (implementation error): %w", err)
		}
		sum += parsedVals[0] * parsedVals[1]
	}
	return sum, nil
}
