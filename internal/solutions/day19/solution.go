package day19

import (
	"fmt"
	"strings"
)

type Solution struct {
	Towels  []string
	Designs []string
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 19:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Possible designs: %d\n", len(PossibleDesigns(s.Designs, s.Towels)))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// PossibleDesigns returns the set of all designs that have one or more solutions with these towels.
func PossibleDesigns(designs []string, towels []string) (possibleDesigns []string) {
	possibleDesigns = make([]string, 0, len(designs))
	for _, design := range designs {
		if isPossible(design, towels) {
			possibleDesigns = append(possibleDesigns, design)
		}
	}

	return possibleDesigns
}

func isPossible(design string, towels []string) bool {
	if len(design) == 0 {
		return true
	}
	for _, towel := range towels {
		if strings.HasPrefix(design, towel) && isPossible(design[len(towel):], towels) {
			return true // ok: towel meets design
		}
	}

	return false
}
