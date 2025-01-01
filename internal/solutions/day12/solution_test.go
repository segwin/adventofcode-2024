package day12_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day12"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day12.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestGetRegions(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		garden day12.Garden

		// outputs
		expected []day12.Region
	}{
		"day's short example 1": {
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: []day12.Region{
				{Kind: 'A', Area: 4, Perimeter: 10},
				{Kind: 'B', Area: 4, Perimeter: 8},
				{Kind: 'C', Area: 4, Perimeter: 10},
				{Kind: 'D', Area: 1, Perimeter: 4},
				{Kind: 'E', Area: 3, Perimeter: 8},
			},
		},
		"day's short example 2": {
			garden: [][]byte{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			},
			expected: []day12.Region{
				{Kind: 'O', Area: 21, Perimeter: 36},
				{Kind: 'X', Area: 1, Perimeter: 4},
				{Kind: 'X', Area: 1, Perimeter: 4},
				{Kind: 'X', Area: 1, Perimeter: 4},
				{Kind: 'X', Area: 1, Perimeter: 4},
			},
		},
		"day's larger example": {
			garden: [][]byte{
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
				{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
				{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
				{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
				{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
				{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
				{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
			},
			expected: []day12.Region{
				{Kind: 'R', Area: 12, Perimeter: 18},
				{Kind: 'I', Area: 4, Perimeter: 8},
				{Kind: 'C', Area: 14, Perimeter: 28},
				{Kind: 'F', Area: 10, Perimeter: 18},
				{Kind: 'V', Area: 13, Perimeter: 20},
				{Kind: 'J', Area: 11, Perimeter: 20},
				{Kind: 'C', Area: 1, Perimeter: 4},
				{Kind: 'E', Area: 13, Perimeter: 18},
				{Kind: 'I', Area: 14, Perimeter: 22},
				{Kind: 'M', Area: 5, Perimeter: 12},
				{Kind: 'S', Area: 3, Perimeter: 8},
			},
		},

		"part 2 example: AB": {
			garden: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			expected: []day12.Region{
				{Kind: 'A', Area: 28, Perimeter: 40},
				{Kind: 'B', Area: 4, Perimeter: 8},
				{Kind: 'B', Area: 4, Perimeter: 8},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day12.GetRegions(tt.garden)
			require.Empty(t, cmp.Diff(
				tt.expected, got,
				cmpopts.SortSlices(lessRegion),
				cmpopts.IgnoreFields(day12.Region{}, "Positions"),
			))
		})
	}
}

func TestTotalFencePrice(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		garden day12.Garden

		// outputs
		expected int
	}{
		"day's short example 1": {
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 140,
		},
		"day's short example 2": {
			garden: [][]byte{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			},
			expected: 772,
		},
		"day's larger example": {
			garden: [][]byte{
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
				{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
				{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
				{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
				{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
				{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
				{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
			},
			expected: 1930,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day12.TotalFencePrice(day12.GetRegions(tt.garden))
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestTotalDiscountedFencePrice(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		garden day12.Garden

		// outputs
		expected int
	}{
		"day's short example 1": {
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 80,
		},
		"day's short example 2": {
			garden: [][]byte{
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
				{'O', 'X', 'O', 'X', 'O'},
				{'O', 'O', 'O', 'O', 'O'},
			},
			expected: 436,
		},
		"day's larger example": {
			garden: [][]byte{
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'F', 'F'},
				{'R', 'R', 'R', 'R', 'I', 'I', 'C', 'C', 'C', 'F'},
				{'V', 'V', 'R', 'R', 'R', 'C', 'C', 'F', 'F', 'F'},
				{'V', 'V', 'R', 'C', 'C', 'C', 'J', 'F', 'F', 'F'},
				{'V', 'V', 'V', 'V', 'C', 'J', 'J', 'C', 'F', 'E'},
				{'V', 'V', 'I', 'V', 'C', 'C', 'J', 'J', 'E', 'E'},
				{'V', 'V', 'I', 'I', 'I', 'C', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'I', 'I', 'J', 'J', 'E', 'E'},
				{'M', 'I', 'I', 'I', 'S', 'I', 'J', 'E', 'E', 'E'},
				{'M', 'M', 'M', 'I', 'S', 'S', 'J', 'E', 'E', 'E'},
			},
			expected: 1206,
		},

		"part 2 example: EX": {
			garden: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			expected: 236,
		},
		"part 2 example: AB": {
			garden: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			expected: 368,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day12.TotalDiscountedFencePrice(day12.GetRegions(tt.garden), tt.garden)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func lessRegion(a, b day12.Region) bool {
	if a.Kind < b.Kind {
		return true
	} else if b.Kind < a.Kind {
		return false
	}
	if a.Area < b.Area {
		return true
	} else if b.Area < a.Area {
		return false
	}
	return a.Perimeter < b.Perimeter
}
