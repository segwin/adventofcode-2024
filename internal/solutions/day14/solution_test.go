package day14_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day14"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day14.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestAfter(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		seconds       int
		initialStates []day14.RobotState
		layout        day14.Layout

		// outputs
		expected []day14.RobotState
	}{
		"day's single-bot example": {
			seconds: 5,
			initialStates: []day14.RobotState{
				{Position: map2d.Position{X: 2, Y: 4}, Velocity: day14.Velocity{X: 2, Y: -3}},
			},
			layout: day14.NewLayout(11, 7),
			expected: []day14.RobotState{
				{Position: map2d.Position{X: 1, Y: 3}, Velocity: day14.Velocity{X: 2, Y: -3}},
			},
		},
		"day's full example": {
			seconds: 100,
			initialStates: []day14.RobotState{
				{Position: map2d.Position{X: 0, Y: 4}, Velocity: day14.Velocity{X: 3, Y: -3}},
				{Position: map2d.Position{X: 6, Y: 3}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 10, Y: 3}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 2, Y: 0}, Velocity: day14.Velocity{X: 2, Y: -1}},
				{Position: map2d.Position{X: 0, Y: 0}, Velocity: day14.Velocity{X: 1, Y: 3}},
				{Position: map2d.Position{X: 3, Y: 0}, Velocity: day14.Velocity{X: -2, Y: -2}},
				{Position: map2d.Position{X: 7, Y: 6}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 3, Y: 0}, Velocity: day14.Velocity{X: -1, Y: -2}},
				{Position: map2d.Position{X: 9, Y: 3}, Velocity: day14.Velocity{X: 2, Y: 3}},
				{Position: map2d.Position{X: 7, Y: 3}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 2, Y: 4}, Velocity: day14.Velocity{X: 2, Y: -3}},
				{Position: map2d.Position{X: 9, Y: 5}, Velocity: day14.Velocity{X: -3, Y: -3}},
			},
			layout: day14.NewLayout(11, 7),
			expected: []day14.RobotState{
				{Position: map2d.Position{X: 3, Y: 5}, Velocity: day14.Velocity{X: 3, Y: -3}},
				{Position: map2d.Position{X: 5, Y: 4}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 9, Y: 0}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 4, Y: 5}, Velocity: day14.Velocity{X: 2, Y: -1}},
				{Position: map2d.Position{X: 1, Y: 6}, Velocity: day14.Velocity{X: 1, Y: 3}},
				{Position: map2d.Position{X: 1, Y: 3}, Velocity: day14.Velocity{X: -2, Y: -2}},
				{Position: map2d.Position{X: 6, Y: 0}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 2, Y: 3}, Velocity: day14.Velocity{X: -1, Y: -2}},
				{Position: map2d.Position{X: 0, Y: 2}, Velocity: day14.Velocity{X: 2, Y: 3}},
				{Position: map2d.Position{X: 6, Y: 0}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 4, Y: 5}, Velocity: day14.Velocity{X: 2, Y: -3}},
				{Position: map2d.Position{X: 6, Y: 6}, Velocity: day14.Velocity{X: -3, Y: -3}},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day14.After(tt.seconds, tt.initialStates, tt.layout)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestSafetyFactor(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		robots []day14.RobotState
		layout day14.Layout

		// outputs
		expected int
	}{
		"day's full example": {
			robots: day14.After(100, []day14.RobotState{
				{Position: map2d.Position{X: 0, Y: 4}, Velocity: day14.Velocity{X: 3, Y: -3}},
				{Position: map2d.Position{X: 6, Y: 3}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 10, Y: 3}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 2, Y: 0}, Velocity: day14.Velocity{X: 2, Y: -1}},
				{Position: map2d.Position{X: 0, Y: 0}, Velocity: day14.Velocity{X: 1, Y: 3}},
				{Position: map2d.Position{X: 3, Y: 0}, Velocity: day14.Velocity{X: -2, Y: -2}},
				{Position: map2d.Position{X: 7, Y: 6}, Velocity: day14.Velocity{X: -1, Y: -3}},
				{Position: map2d.Position{X: 3, Y: 0}, Velocity: day14.Velocity{X: -1, Y: -2}},
				{Position: map2d.Position{X: 9, Y: 3}, Velocity: day14.Velocity{X: 2, Y: 3}},
				{Position: map2d.Position{X: 7, Y: 3}, Velocity: day14.Velocity{X: -1, Y: 2}},
				{Position: map2d.Position{X: 2, Y: 4}, Velocity: day14.Velocity{X: 2, Y: -3}},
				{Position: map2d.Position{X: 9, Y: 5}, Velocity: day14.Velocity{X: -3, Y: -3}},
			}, day14.NewLayout(11, 7)),
			layout:   day14.NewLayout(11, 7),
			expected: 12,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day14.SafetyFactor(tt.robots, tt.layout)
			assert.Equal(t, tt.expected, got)
		})
	}
}
