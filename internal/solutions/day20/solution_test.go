package day20_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day20"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day20.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestNavigate(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		racetrack  map2d.Map[day20.Tile]
		start, end map2d.Position

		// outputs
		expected []map2d.Position
	}{
		"day's example": {
			racetrack: map2d.Map[day20.Tile]{
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', 'S', '#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '.', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '#', '#', '.', '.', 'E', '#', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '.', '#', '#', '#', '#', '#', '#', '#', '.', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '#', '#', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '.', '#', '#', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', '.', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
			},
			start: map2d.Position{X: 1, Y: 3},
			end:   map2d.Position{X: 5, Y: 7},
			expected: []map2d.Position{
				{X: 1, Y: 3}, {X: 1, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1},
				{X: 3, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 5, Y: 2},
				{X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 7, Y: 2}, {X: 7, Y: 3},
				{X: 7, Y: 4}, {X: 7, Y: 5}, {X: 7, Y: 6}, {X: 7, Y: 7}, {X: 8, Y: 7},
				{X: 9, Y: 7}, {X: 9, Y: 6}, {X: 9, Y: 5}, {X: 9, Y: 4}, {X: 9, Y: 3},
				{X: 9, Y: 2}, {X: 9, Y: 1}, {X: 10, Y: 1}, {X: 11, Y: 1}, {X: 12, Y: 1},
				{X: 13, Y: 1}, {X: 13, Y: 2}, {X: 13, Y: 3}, {X: 12, Y: 3}, {X: 11, Y: 3},
				{X: 11, Y: 4}, {X: 11, Y: 5}, {X: 12, Y: 5}, {X: 13, Y: 5}, {X: 13, Y: 6},
				{X: 13, Y: 7}, {X: 12, Y: 7}, {X: 11, Y: 7}, {X: 11, Y: 8}, {X: 11, Y: 9},
				{X: 12, Y: 9}, {X: 13, Y: 9}, {X: 13, Y: 10}, {X: 13, Y: 11}, {X: 12, Y: 11},
				{X: 11, Y: 11}, {X: 11, Y: 12}, {X: 11, Y: 13}, {X: 10, Y: 13}, {X: 9, Y: 13},
				{X: 9, Y: 12}, {X: 9, Y: 11}, {X: 9, Y: 10}, {X: 9, Y: 9}, {X: 8, Y: 9},
				{X: 7, Y: 9}, {X: 7, Y: 10}, {X: 7, Y: 11}, {X: 7, Y: 12}, {X: 7, Y: 13},
				{X: 6, Y: 13}, {X: 5, Y: 13}, {X: 5, Y: 12}, {X: 5, Y: 11}, {X: 4, Y: 11},
				{X: 3, Y: 11}, {X: 3, Y: 12}, {X: 3, Y: 13}, {X: 2, Y: 13}, {X: 1, Y: 13},
				{X: 1, Y: 12}, {X: 1, Y: 11}, {X: 1, Y: 10}, {X: 1, Y: 9}, {X: 2, Y: 9},
				{X: 3, Y: 9}, {X: 3, Y: 8}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day20.Navigate(tt.racetrack, tt.start, tt.end)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestFindCheats(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		racetrack   map2d.Map[day20.Tile]
		normalRoute []map2d.Position
		minSavings  int

		// outputs
		expected map[day20.Cheat]int
	}{
		"day's example": {
			racetrack: map2d.Map[day20.Tile]{
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', 'S', '#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '.', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '#', '#', '.', '.', 'E', '#', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '#', '#', '.', '#', '#', '#', '#', '#', '#', '#', '.', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '#', '#', '.', '.', '.', '#', '.', '.', '.', '#'},
				{'#', '.', '#', '#', '#', '#', '#', '.', '#', '.', '#', '#', '#', '.', '#'},
				{'#', '.', '#', '.', '.', '.', '#', '.', '#', '.', '#', '.', '.', '.', '#'},
				{'#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '.', '#', '#', '#'},
				{'#', '.', '.', '.', '#', '.', '.', '.', '#', '.', '.', '.', '#', '#', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
			},
			normalRoute: []map2d.Position{
				{X: 1, Y: 3}, {X: 1, Y: 2}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 3, Y: 1},
				{X: 3, Y: 2}, {X: 3, Y: 3}, {X: 4, Y: 3}, {X: 5, Y: 3}, {X: 5, Y: 2},
				{X: 5, Y: 1}, {X: 6, Y: 1}, {X: 7, Y: 1}, {X: 7, Y: 2}, {X: 7, Y: 3},
				{X: 7, Y: 4}, {X: 7, Y: 5}, {X: 7, Y: 6}, {X: 7, Y: 7}, {X: 8, Y: 7},
				{X: 9, Y: 7}, {X: 9, Y: 6}, {X: 9, Y: 5}, {X: 9, Y: 4}, {X: 9, Y: 3},
				{X: 9, Y: 2}, {X: 9, Y: 1}, {X: 10, Y: 1}, {X: 11, Y: 1}, {X: 12, Y: 1},
				{X: 13, Y: 1}, {X: 13, Y: 2}, {X: 13, Y: 3}, {X: 12, Y: 3}, {X: 11, Y: 3},
				{X: 11, Y: 4}, {X: 11, Y: 5}, {X: 12, Y: 5}, {X: 13, Y: 5}, {X: 13, Y: 6},
				{X: 13, Y: 7}, {X: 12, Y: 7}, {X: 11, Y: 7}, {X: 11, Y: 8}, {X: 11, Y: 9},
				{X: 12, Y: 9}, {X: 13, Y: 9}, {X: 13, Y: 10}, {X: 13, Y: 11}, {X: 12, Y: 11},
				{X: 11, Y: 11}, {X: 11, Y: 12}, {X: 11, Y: 13}, {X: 10, Y: 13}, {X: 9, Y: 13},
				{X: 9, Y: 12}, {X: 9, Y: 11}, {X: 9, Y: 10}, {X: 9, Y: 9}, {X: 8, Y: 9},
				{X: 7, Y: 9}, {X: 7, Y: 10}, {X: 7, Y: 11}, {X: 7, Y: 12}, {X: 7, Y: 13},
				{X: 6, Y: 13}, {X: 5, Y: 13}, {X: 5, Y: 12}, {X: 5, Y: 11}, {X: 4, Y: 11},
				{X: 3, Y: 11}, {X: 3, Y: 12}, {X: 3, Y: 13}, {X: 2, Y: 13}, {X: 1, Y: 13},
				{X: 1, Y: 12}, {X: 1, Y: 11}, {X: 1, Y: 10}, {X: 1, Y: 9}, {X: 2, Y: 9},
				{X: 3, Y: 9}, {X: 3, Y: 8}, {X: 3, Y: 7}, {X: 4, Y: 7}, {X: 5, Y: 7},
			},
			minSavings: 0,
			expected: map[day20.Cheat]int{
				{Start: map2d.Position{X: 2, Y: 2}, End: map2d.Position{X: 3, Y: 2}}:     2,
				{Start: map2d.Position{X: 2, Y: 3}, End: map2d.Position{X: 3, Y: 3}}:     4,
				{Start: map2d.Position{X: 4, Y: 1}, End: map2d.Position{X: 5, Y: 1}}:     4,
				{Start: map2d.Position{X: 4, Y: 2}, End: map2d.Position{X: 5, Y: 2}}:     2,
				{Start: map2d.Position{X: 6, Y: 3}, End: map2d.Position{X: 7, Y: 3}}:     4,
				{Start: map2d.Position{X: 6, Y: 2}, End: map2d.Position{X: 7, Y: 2}}:     2,
				{Start: map2d.Position{X: 8, Y: 1}, End: map2d.Position{X: 9, Y: 1}}:     12,
				{Start: map2d.Position{X: 8, Y: 2}, End: map2d.Position{X: 9, Y: 2}}:     10,
				{Start: map2d.Position{X: 8, Y: 3}, End: map2d.Position{X: 9, Y: 3}}:     8,
				{Start: map2d.Position{X: 8, Y: 4}, End: map2d.Position{X: 9, Y: 4}}:     6,
				{Start: map2d.Position{X: 8, Y: 5}, End: map2d.Position{X: 9, Y: 5}}:     4,
				{Start: map2d.Position{X: 8, Y: 6}, End: map2d.Position{X: 9, Y: 6}}:     2,
				{Start: map2d.Position{X: 6, Y: 7}, End: map2d.Position{X: 5, Y: 7}}:     64,
				{Start: map2d.Position{X: 7, Y: 8}, End: map2d.Position{X: 7, Y: 9}}:     40,
				{Start: map2d.Position{X: 8, Y: 8}, End: map2d.Position{X: 8, Y: 9}}:     38,
				{Start: map2d.Position{X: 9, Y: 8}, End: map2d.Position{X: 9, Y: 9}}:     36,
				{Start: map2d.Position{X: 10, Y: 7}, End: map2d.Position{X: 11, Y: 7}}:   20,
				{Start: map2d.Position{X: 10, Y: 5}, End: map2d.Position{X: 11, Y: 5}}:   12,
				{Start: map2d.Position{X: 10, Y: 4}, End: map2d.Position{X: 11, Y: 4}}:   10,
				{Start: map2d.Position{X: 10, Y: 3}, End: map2d.Position{X: 11, Y: 3}}:   8,
				{Start: map2d.Position{X: 11, Y: 2}, End: map2d.Position{X: 11, Y: 3}}:   4,
				{Start: map2d.Position{X: 12, Y: 2}, End: map2d.Position{X: 12, Y: 3}}:   2,
				{Start: map2d.Position{X: 13, Y: 4}, End: map2d.Position{X: 13, Y: 5}}:   4,
				{Start: map2d.Position{X: 12, Y: 4}, End: map2d.Position{X: 12, Y: 5}}:   2,
				{Start: map2d.Position{X: 11, Y: 6}, End: map2d.Position{X: 11, Y: 7}}:   4,
				{Start: map2d.Position{X: 12, Y: 6}, End: map2d.Position{X: 12, Y: 7}}:   2,
				{Start: map2d.Position{X: 13, Y: 8}, End: map2d.Position{X: 13, Y: 9}}:   4,
				{Start: map2d.Position{X: 12, Y: 8}, End: map2d.Position{X: 12, Y: 9}}:   2,
				{Start: map2d.Position{X: 10, Y: 9}, End: map2d.Position{X: 9, Y: 9}}:    12,
				{Start: map2d.Position{X: 11, Y: 10}, End: map2d.Position{X: 11, Y: 11}}: 4,
				{Start: map2d.Position{X: 12, Y: 10}, End: map2d.Position{X: 12, Y: 11}}: 2,
				{Start: map2d.Position{X: 10, Y: 11}, End: map2d.Position{X: 9, Y: 11}}:  4,
				{Start: map2d.Position{X: 10, Y: 12}, End: map2d.Position{X: 9, Y: 12}}:  2,
				{Start: map2d.Position{X: 8, Y: 13}, End: map2d.Position{X: 7, Y: 13}}:   8,
				{Start: map2d.Position{X: 8, Y: 12}, End: map2d.Position{X: 7, Y: 12}}:   6,
				{Start: map2d.Position{X: 8, Y: 11}, End: map2d.Position{X: 7, Y: 11}}:   4,
				{Start: map2d.Position{X: 8, Y: 10}, End: map2d.Position{X: 7, Y: 10}}:   2,
				{Start: map2d.Position{X: 6, Y: 11}, End: map2d.Position{X: 5, Y: 11}}:   4,
				{Start: map2d.Position{X: 6, Y: 12}, End: map2d.Position{X: 5, Y: 12}}:   2,
				{Start: map2d.Position{X: 4, Y: 13}, End: map2d.Position{X: 3, Y: 13}}:   4,
				{Start: map2d.Position{X: 4, Y: 12}, End: map2d.Position{X: 3, Y: 12}}:   2,
				{Start: map2d.Position{X: 3, Y: 10}, End: map2d.Position{X: 3, Y: 9}}:    8,
				{Start: map2d.Position{X: 2, Y: 11}, End: map2d.Position{X: 1, Y: 11}}:   4,
				{Start: map2d.Position{X: 2, Y: 12}, End: map2d.Position{X: 1, Y: 12}}:   2,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day20.FindCheats(tt.racetrack, tt.normalRoute, tt.minSavings)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}
