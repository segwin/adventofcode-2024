package day11

import "math"

type Stone int

func (s Stone) Change() []Stone {
	// 0 => 1
	if s == 0 {
		return []Stone{1}
	}

	// even number of digits => split into 2
	if pow10 := int(math.Log10(float64(s))) + 1; pow10%2 == 0 {
		midDigits := math.Pow10(pow10 / 2)
		return []Stone{s / Stone(midDigits), s % Stone(midDigits)}
	}

	// everything else => multiply by 2024
	return []Stone{2024 * s}
}

// stoneCounts stores a single copy of each stone to reduce the number of calculations needed.
type stoneCounts map[Stone]int

func (sc stoneCounts) Blink() stoneCounts {
	newCounts := make(map[Stone]int, len(sc))
	for stone, count := range sc {
		for _, newStone := range stone.Change() {
			newCounts[newStone] += count
		}
	}
	return newCounts
}

func (sc stoneCounts) Total() int {
	sum := 0
	for _, count := range sc {
		sum += count
	}
	return sum
}
