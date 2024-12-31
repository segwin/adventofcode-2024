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
		"day's simplified example": {
			terrain: [][]int{
				{0, 1, 2, 3},
				{1, 2, 3, 4},
				{8, 7, 6, 5},
				{9, 8, 7, 6},
			},
			expected: []day10.Trailhead{
				{Position: map2d.Position{X: 0, Y: 0}, Score: 1},
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
				{Position: map2d.Position{X: 0, Y: 6}, Score: 5},
				{Position: map2d.Position{X: 1, Y: 7}, Score: 5},
				{Position: map2d.Position{X: 2, Y: 0}, Score: 5},
				{Position: map2d.Position{X: 2, Y: 5}, Score: 1},
				{Position: map2d.Position{X: 4, Y: 0}, Score: 6},
				{Position: map2d.Position{X: 4, Y: 2}, Score: 5},
				{Position: map2d.Position{X: 5, Y: 5}, Score: 3},
				{Position: map2d.Position{X: 6, Y: 4}, Score: 3},
				{Position: map2d.Position{X: 6, Y: 6}, Score: 3},
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
