package day12_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day12"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
)

func TestRegion_Sides(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		region day12.Region
		garden day12.Garden

		// outputs
		expected int
	}{
		"ABCDE: region A": {
			region: day12.Region{Kind: 'A', Positions: toPositionSet([]map2d.Position{
				{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 0, X: 2}, {Y: 0, X: 3},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 4,
		},
		"ABCDE: region B": {
			region: day12.Region{Kind: 'B', Positions: toPositionSet([]map2d.Position{
				{Y: 1, X: 0}, {Y: 1, X: 1},
				{Y: 2, X: 0}, {Y: 2, X: 1},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 4,
		},
		"ABCDE: region C": {
			region: day12.Region{Kind: 'C', Positions: toPositionSet([]map2d.Position{
				{Y: 1, X: 2}, {Y: 2, X: 2},
				{Y: 2, X: 3}, {Y: 3, X: 3},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 8,
		},
		"ABCDE: region D": {
			region: day12.Region{Kind: 'D', Positions: toPositionSet([]map2d.Position{
				{Y: 1, X: 3},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 4,
		},
		"ABCDE: region E": {
			region: day12.Region{Kind: 'E', Positions: toPositionSet([]map2d.Position{
				{Y: 3, X: 0}, {Y: 3, X: 1}, {Y: 3, X: 2},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'B', 'B', 'C', 'D'},
				{'B', 'B', 'C', 'C'},
				{'E', 'E', 'E', 'C'},
			},
			expected: 4,
		},

		"EX: region E": {
			region: day12.Region{Kind: 'E', Positions: toPositionSet([]map2d.Position{
				{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 0, X: 2}, {Y: 0, X: 3}, {Y: 0, X: 4},
				{Y: 1, X: 0},
				{Y: 2, X: 0}, {Y: 2, X: 1}, {Y: 2, X: 2}, {Y: 2, X: 3}, {Y: 2, X: 4},
				{Y: 3, X: 0},
				{Y: 4, X: 0}, {Y: 4, X: 1}, {Y: 4, X: 2}, {Y: 4, X: 3}, {Y: 4, X: 4},
			})},
			garden: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			expected: 12,
		},
		"EX: region X1": {
			region: day12.Region{Kind: 'X', Positions: toPositionSet([]map2d.Position{
				{Y: 1, X: 1}, {Y: 1, X: 2}, {Y: 1, X: 3}, {Y: 1, X: 4},
			})},
			garden: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			expected: 4,
		},
		"EX: region X2": {
			region: day12.Region{Kind: 'X', Positions: toPositionSet([]map2d.Position{
				{Y: 3, X: 1}, {Y: 3, X: 2}, {Y: 3, X: 3}, {Y: 3, X: 4},
			})},
			garden: [][]byte{
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
				{'E', 'X', 'X', 'X', 'X'},
				{'E', 'E', 'E', 'E', 'E'},
			},
			expected: 4,
		},

		"AB: region A": {
			region: day12.Region{Kind: 'A', Positions: toPositionSet([]map2d.Position{
				{Y: 0, X: 0}, {Y: 0, X: 1}, {Y: 0, X: 2}, {Y: 0, X: 3}, {Y: 0, X: 4}, {Y: 0, X: 5},
				{Y: 1, X: 0}, {Y: 1, X: 1}, {Y: 1, X: 2}, {Y: 1, X: 5},
				{Y: 2, X: 0}, {Y: 2, X: 1}, {Y: 2, X: 2}, {Y: 2, X: 5},
				{Y: 3, X: 0}, {Y: 3, X: 3}, {Y: 3, X: 4}, {Y: 3, X: 5},
				{Y: 4, X: 0}, {Y: 4, X: 3}, {Y: 4, X: 4}, {Y: 4, X: 5},
				{Y: 5, X: 0}, {Y: 5, X: 1}, {Y: 5, X: 2}, {Y: 5, X: 3}, {Y: 5, X: 4}, {Y: 5, X: 5},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			expected: 12,
		},
		"AB: region B1": {
			region: day12.Region{Kind: 'A', Positions: toPositionSet([]map2d.Position{
				{Y: 1, X: 3}, {Y: 1, X: 4},
				{Y: 2, X: 3}, {Y: 2, X: 4},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			expected: 4,
		},
		"AB: region B2": {
			region: day12.Region{Kind: 'A', Positions: toPositionSet([]map2d.Position{
				{Y: 3, X: 1}, {Y: 3, X: 2},
				{Y: 4, X: 1}, {Y: 4, X: 2},
			})},
			garden: [][]byte{
				{'A', 'A', 'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'A', 'A', 'B', 'B', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'B', 'B', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A', 'A', 'A'},
			},
			expected: 4,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.region.Sides(tt.garden)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func toPositionSet(positions []map2d.Position) map[map2d.Position]struct{} {
	out := make(map[map2d.Position]struct{}, len(positions))
	for _, pos := range positions {
		out[pos] = struct{}{}
	}
	return out
}
