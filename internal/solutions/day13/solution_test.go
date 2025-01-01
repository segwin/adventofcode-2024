package day13_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day13"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day13.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestCost(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		a int
		b int

		// outputs
		expected int
	}{
		"no presses":                {a: 0, b: 0, expected: 0},
		"a=1, b=0":                  {a: 1, b: 0, expected: day13.ACost},
		"a=0, b=1":                  {a: 0, b: 1, expected: day13.BCost},
		"day's example: a=80, b=40": {a: 80, b: 40, expected: 280},
		"day's example: a=38, b=86": {a: 38, b: 86, expected: 200},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day13.Cost(tt.a, tt.b)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestOptimalPresses(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		m day13.ClawMachine

		// outputs
		expectedA  int
		expectedB  int
		expectedOK bool
	}{
		"ok: day's 1st example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 94, Y: 34},
				MoveB: map2d.Distance{X: 22, Y: 67},
				Prize: map2d.Distance{X: 8400, Y: 5400},
			},
			expectedA: 80, expectedB: 40, expectedOK: true,
		},
		"not ok: day's 2nd example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 26, Y: 66},
				MoveB: map2d.Distance{X: 67, Y: 21},
				Prize: map2d.Distance{X: 12748, Y: 12176},
			},
			expectedOK: false,
		},
		"not ok: day's 3rd example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 17, Y: 86},
				MoveB: map2d.Distance{X: 84, Y: 37},
				Prize: map2d.Distance{X: 7870, Y: 6450},
			},
			expectedA: 38, expectedB: 86, expectedOK: true,
		},
		"ok: day's 4th example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 69, Y: 23},
				MoveB: map2d.Distance{X: 27, Y: 71},
				Prize: map2d.Distance{X: 18641, Y: 10279},
			},
			expectedOK: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			a, b, ok := day13.OptimalPresses(tt.m)
			require.Equal(t, tt.expectedOK, ok)
			assert.Equal(t, tt.expectedA, a)
			assert.Equal(t, tt.expectedB, b)
		})
	}
}
