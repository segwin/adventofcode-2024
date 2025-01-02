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
		"day's example: a=118679050709, b=103199174542": {a: 118679050709, b: 103199174542, expected: 459236326669},
		"day's example: a=102851800151, b=107526881786": {a: 102851800151, b: 107526881786, expected: 416082282239},
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
				Prize: map2d.Position{X: 8400, Y: 5400},
			},
			expectedA: 80, expectedB: 40, expectedOK: true,
		},
		"not ok: day's 2nd example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 26, Y: 66},
				MoveB: map2d.Distance{X: 67, Y: 21},
				Prize: map2d.Position{X: 12748, Y: 12176},
			},
			expectedOK: false,
		},
		"not ok: day's 3rd example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 17, Y: 86},
				MoveB: map2d.Distance{X: 84, Y: 37},
				Prize: map2d.Position{X: 7870, Y: 6450},
			},
			expectedA: 38, expectedB: 86, expectedOK: true,
		},
		"ok: day's 4th example": {
			m: day13.ClawMachine{
				MoveA: map2d.Distance{X: 69, Y: 23},
				MoveB: map2d.Distance{X: 27, Y: 71},
				Prize: map2d.Position{X: 18641, Y: 10279},
			},
			expectedOK: false,
		},

		"not ok: day's 1st example (adjusted, part 2)": {
			m: correctOne(day13.ClawMachine{
				MoveA: map2d.Distance{X: 94, Y: 34},
				MoveB: map2d.Distance{X: 22, Y: 67},
				Prize: map2d.Position{X: 8400, Y: 5400},
			}),
			expectedOK: false,
		},
		"ok: day's 2nd example (adjusted, part 2)": {
			m: correctOne(day13.ClawMachine{
				MoveA: map2d.Distance{X: 26, Y: 66},
				MoveB: map2d.Distance{X: 67, Y: 21},
				Prize: map2d.Position{X: 12748, Y: 12176},
			}),
			expectedA: 118679050709, expectedB: 103199174542, expectedOK: true,
		},
		"not ok: day's 3rd example (adjusted, part 2)": {
			m: correctOne(day13.ClawMachine{
				MoveA: map2d.Distance{X: 17, Y: 86},
				MoveB: map2d.Distance{X: 84, Y: 37},
				Prize: map2d.Position{X: 7870, Y: 6450},
			}),
			expectedOK: false,
		},
		"ok: day's 4th example (adjusted, part 2)": {
			m: correctOne(day13.ClawMachine{
				MoveA: map2d.Distance{X: 69, Y: 23},
				MoveB: map2d.Distance{X: 27, Y: 71},
				Prize: map2d.Position{X: 18641, Y: 10279},
			}),
			expectedA: 102851800151, expectedB: 107526881786, expectedOK: true,
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

func correctOne(m day13.ClawMachine) day13.ClawMachine {
	return day13.CorrectMachines(1e13, m)[0]
}
