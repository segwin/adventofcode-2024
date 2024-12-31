package day11

import "fmt"

type Solution struct {
	InitialStones []Stone
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 11:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total stones after 25 blinks: %d\n", len(BlinkTimes(s.InitialStones, 25)))

	return nil
}

func Blink(stones []Stone) []Stone {
	newStones := make([]Stone, 0, len(stones))
	for _, s := range stones {
		newStones = append(newStones, s.Change()...)
	}
	return newStones
}

func BlinkTimes(stones []Stone, times int) []Stone {
	for range times {
		stones = Blink(stones)
	}
	return stones
}
