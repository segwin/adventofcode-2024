package day19_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day19"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day19.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestIsPossible(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		designs []day19.Design
		towels  []day19.Towel

		// outputs
		expected int
	}{
		"day's example": {
			designs:  []day19.Design{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"},
			towels:   []day19.Towel{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			expected: 6,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day19.PossibleDesigns(tt.designs, tt.towels)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCombinationsByDesign(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		designs []day19.Design
		towels  []day19.Towel

		// outputs
		expected int
	}{
		"day's example": {
			designs:  []day19.Design{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"},
			towels:   []day19.Towel{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			expected: 16,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day19.TotalCombinations(tt.designs, tt.towels)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func BenchmarkCombinationsByDesign(b *testing.B) {
	for range b.N {
		day19.TotalCombinations([]day19.Design{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"}, []day19.Towel{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"})
	}
}
