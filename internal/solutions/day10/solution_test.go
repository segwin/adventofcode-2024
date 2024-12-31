package day10_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day10"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
)

func TestGetTrailheads(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		terrain day10.TopographicMap

		// outputs
		expected []day10.Trailhead
	}{
		"day's small example": {
			terrain: [][]int{
				{0, 1, 2, 3},
				{1, 2, 3, 4},
				{8, 7, 6, 5},
				{9, 8, 7, 6},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{X: 0, Y: 0}, Score: 1, Rating: 16},
			},
		},
		"day's full example": {
			terrain: [][]int{
				{8, 9, 0, 1, 0, 1, 2, 3},
				{7, 8, 1, 2, 1, 8, 7, 4},
				{8, 7, 4, 3, 0, 9, 6, 5},
				{9, 6, 5, 4, 9, 8, 7, 4},
				{4, 5, 6, 7, 8, 9, 0, 3},
				{3, 2, 0, 1, 9, 0, 1, 2},
				{0, 1, 3, 2, 9, 8, 0, 1},
				{1, 0, 4, 5, 6, 7, 3, 2},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{Y: 0, X: 2}, Score: 5, Rating: 20},
				{Position: map2d.Position{Y: 0, X: 4}, Score: 6, Rating: 24},
				{Position: map2d.Position{Y: 2, X: 4}, Score: 5, Rating: 10},
				{Position: map2d.Position{Y: 4, X: 6}, Score: 3, Rating: 4},
				{Position: map2d.Position{Y: 5, X: 2}, Score: 1, Rating: 1},
				{Position: map2d.Position{Y: 5, X: 5}, Score: 3, Rating: 4},
				{Position: map2d.Position{Y: 6, X: 0}, Score: 5, Rating: 5},
				{Position: map2d.Position{Y: 6, X: 6}, Score: 3, Rating: 8},
				{Position: map2d.Position{Y: 7, X: 1}, Score: 5, Rating: 5},
			},
		},

		"day's simplified example 1": {
			terrain: [][]int{
				{-1, -1, -1, -1, -1, 0, -1},
				{-1, -1, 4, 3, 2, 1, -1},
				{-1, -1, 5, -1, -1, 2, -1},
				{-1, -1, 6, 5, 4, 3, -1},
				{-1, -1, 7, -1, -1, 4, -1},
				{-1, -1, 8, 7, 6, 5, -1},
				{-1, -1, 9, -1, -1, -1, -1},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{X: 5, Y: 0}, Score: 1, Rating: 3},
			},
		},
		"day's simplified example 2": {
			terrain: [][]int{
				{-1, -1, 9, 0, -1, -1, 9},
				{-1, -1, -1, 1, -1, 9, 8},
				{-1, -1, -1, 2, -1, -1, 7},
				{6, 5, 4, 3, 4, 5, 6},
				{7, 6, 5, -1, 9, 8, 7},
				{8, 7, 6, -1, -1, -1, -1},
				{9, 8, 7, -1, -1, -1, -1},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{X: 3, Y: 0}, Score: 4, Rating: 13},
			},
		},
		"day's simplified example 3": {
			terrain: [][]int{
				{0, 1, 2, 3, 4, 5},
				{1, 2, 3, 4, 5, 6},
				{2, 3, 4, 5, 6, 7},
				{3, 4, 5, 6, 7, 8},
				{4, -1, 6, 7, 8, 9},
				{5, 6, 7, 8, 9, -1},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{X: 0, Y: 0}, Score: 2, Rating: 227},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day10.GetTrailheads(tt.terrain)
			assert.Empty(t, cmp.Diff(tt.expected, got, cmpopts.SortSlices(func(a, b day10.Trailhead) bool {
				if a.Position.LessThan(b.Position) {
					return true
				} else if b.Position.LessThan(a.Position) {
					return false
				}
				return a.Score < b.Score
			})))
		})
	}
}
