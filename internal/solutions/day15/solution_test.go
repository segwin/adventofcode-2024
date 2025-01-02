package day15_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day15"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day15.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestResolve(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		layout       day15.Layout
		instructions []map2d.Direction

		// outputs
		expected day15.Layout
	}{
		"day's short example": {
			layout: day15.Layout{Map: [][]day15.Tile{
				{'#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', 'O', '.', 'O', '.', '#'},
				{'#', '#', '@', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '.', '#', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#'},
			}},
			instructions: []map2d.Direction{
				// <^^>>
				// >vv<v
				// >>v<<
				map2d.West(), map2d.North(), map2d.North(), map2d.East(), map2d.East(),
				map2d.East(), map2d.South(), map2d.South(), map2d.West(), map2d.South(),
				map2d.East(), map2d.East(), map2d.South(), map2d.West(), map2d.West(),
			},
			expected: day15.Layout{Map: [][]day15.Tile{
				{'#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', '.', '.', 'O', 'O', '#'},
				{'#', '#', '.', '.', '.', '.', '.', '#'},
				{'#', '.', '.', '.', '.', '.', 'O', '#'},
				{'#', '.', '#', 'O', '@', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#'},
			}},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day15.Resolve(tt.layout, tt.instructions...)
			require.NoError(t, err)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}
