package day19_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
		designs []string
		towels  []string

		// outputs
		expected []string
	}{
		"day's example": {
			designs:  []string{"brwrr", "bggr", "gbbr", "rrbgbr", "ubwu", "bwurrg", "brgr", "bbrgwb"},
			towels:   []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"},
			expected: []string{"brwrr", "bggr", "gbbr", "rrbgbr", "bwurrg", "brgr"},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day19.PossibleDesigns(tt.designs, tt.towels)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}
