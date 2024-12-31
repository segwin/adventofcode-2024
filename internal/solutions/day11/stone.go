package day11

import "math"

type Stone int

func (s Stone) Change() []Stone {
	// 0 => 1
	if s == 0 {
		return []Stone{1}
	}

	// even number of digits => split into 2
	if pow10 := int(math.Floor(math.Log10(float64(s)))) + 1; pow10 > 0 && pow10%2 == 0 {
		midDigits := math.Pow10(pow10 / 2)
		return []Stone{s / Stone(midDigits), s % Stone(midDigits)}
	}

	// everything else => multiply by 2024
	return []Stone{2024 * s}
}
