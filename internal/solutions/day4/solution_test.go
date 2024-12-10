package day4_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day4"
	"github.com/stretchr/testify/assert"
)

func TestCountXMAS(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		search []string

		// outputs
		expected int
	}{
		"ok: day's example": {
			search: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 18,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day4.CountXMAS(tt.search)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCountCrossMas(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		search []string

		// outputs
		expected int
	}{
		"ok: day's example": {
			search: []string{
				"MMMSXXMASM",
				"MSAMXMSMSA",
				"AMXSXMAAMM",
				"MSAMASMSMX",
				"XMASAMXAMM",
				"XXAMMXXAMA",
				"SMSMSASXSS",
				"SAXAMASAAA",
				"MAMMMXMMMM",
				"MXMXAXMASX",
			},
			expected: 9,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day4.CountCrossMas(tt.search)
			assert.Equal(t, tt.expected, got)
		})
	}
}
