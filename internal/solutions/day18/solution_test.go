package day18_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day18"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day18.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestDropBytes(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		memory    day18.Layout
		positions []map2d.Position

		// outputs
		expected day18.Layout
	}{
		"day's example": {
			memory: day18.NewEmptyLayout(7, 7),
			positions: []map2d.Position{
				{X: 5, Y: 4}, {X: 4, Y: 2}, {X: 4, Y: 5}, {X: 3, Y: 0}, {X: 2, Y: 1}, {X: 6, Y: 3},
				{X: 2, Y: 4}, {X: 1, Y: 5}, {X: 0, Y: 6}, {X: 3, Y: 3}, {X: 2, Y: 6}, {X: 5, Y: 1},
			},
			expected: day18.Layout{
				{'.', '.', '.', '#', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '.', '#', '.', '.', '#'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '#', '.', '.', '#', '.', '.'},
				{'#', '.', '#', '.', '.', '.', '.'},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day18.DropBytes(tt.memory, tt.positions...)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestSolveDistance(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		memory     day18.Layout
		start, end map2d.Position

		// outputs
		expectedDistance int
	}{
		"day's example": {
			memory: day18.Layout{
				{'.', '.', '.', '#', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '.', '#', '.', '.', '#'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '#', '.', '.', '#', '.', '.'},
				{'#', '.', '#', '.', '.', '.', '.'},
			},
			start:            map2d.Position{X: 0, Y: 0},
			end:              map2d.Position{X: 6, Y: 6},
			expectedDistance: 22,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day18.Solve(tt.memory, tt.start, tt.end)
			assert.Equal(t, tt.expectedDistance, day18.Distance(got))
		})
	}
}

func TestFirstBlocking(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		memory       day18.Layout
		start, end   map2d.Position
		fallingBytes []map2d.Position

		// outputs
		expected map2d.Position
	}{
		"day's example": {
			memory: day18.Layout{
				{'.', '.', '.', '#', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '.', '#', '.', '.', '#'},
				{'.', '.', '#', '.', '.', '#', '.'},
				{'.', '#', '.', '.', '#', '.', '.'},
				{'#', '.', '#', '.', '.', '.', '.'},
			},
			start: map2d.Position{X: 0, Y: 0},
			end:   map2d.Position{X: 6, Y: 6},
			fallingBytes: []map2d.Position{
				// omits first 12 bytes: already included in memory
				{X: 1, Y: 2}, {X: 5, Y: 5}, {X: 2, Y: 5}, {X: 6, Y: 5}, {X: 1, Y: 4},
				{X: 0, Y: 4}, {X: 6, Y: 4}, {X: 1, Y: 1}, {X: 6, Y: 1}, {X: 1, Y: 0},
				{X: 0, Y: 5}, {X: 1, Y: 6}, {X: 2, Y: 0},
			},
			expected: map2d.Position{X: 6, Y: 1},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day18.FirstBlockingByte(tt.memory, tt.start, tt.end, tt.fallingBytes)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}
