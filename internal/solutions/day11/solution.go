package day11

import (
	"fmt"
)

type Solution struct {
	InitialStones []Stone
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 11:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total stones after 25 blinks: %d\n", Blink(s.InitialStones, 25))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Total stones after 75 blinks: %d\n", Blink(s.InitialStones, 75))

	return nil
}

func Blink(stones []Stone, times int) (finalCount int) {
	sc := stoneCounts{}
	for _, stone := range stones {
		sc[stone]++
	}

	for range times {
		sc = sc.Blink()
	}

	return sc.Total()
}
