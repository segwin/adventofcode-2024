package day6_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCountGuardPositions(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		floorMap [][]day6.Tile

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example": {
			floorMap: [][]day6.Tile{
				{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
			},
			expected: 41,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, _, err := day6.CountGuardPositions(tt.floorMap)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCountLoopPositions(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		floorMap [][]day6.Tile

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example": {
			floorMap: [][]day6.Tile{
				{'.', '.', '.', '.', '#', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '#', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '#', '.', '.', '^', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
				{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '#', '.', '.', '.'},
			},
			expected: 6,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			_, originalStates, err := day6.CountGuardPositions(tt.floorMap)
			require.NoError(t, err)

			got, err := day6.CountLoopPositions(tt.floorMap, originalStates)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
