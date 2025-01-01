package day7_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day7.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestTotalCalibration(t *testing.T) {
	t.Parallel()

	exampleInput := map[int][]int{
		190:    {10, 19},
		3267:   {81, 40, 27},
		83:     {17, 5},
		156:    {15, 6},
		7290:   {6, 8, 6, 15},
		161011: {16, 10, 13},
		192:    {17, 8, 14},
		21037:  {9, 7, 18, 13},
		292:    {11, 6, 16, 20},
	}

	tests := map[string]struct {
		// inputs
		operandsByResult map[int][]int
		operators        []day7.Operator

		// outputs
		expected int
	}{
		"ok: day's example, part 1": {
			operandsByResult: exampleInput,
			operators:        []day7.Operator{day7.Plus{}, day7.Star{}},
			expected:         3749,
		},
		"ok: day's example, part 2": {
			operandsByResult: exampleInput,
			operators:        []day7.Operator{day7.Plus{}, day7.Star{}, day7.DoublePipe{}},
			expected:         11387,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day7.TotalCalibration(tt.operandsByResult, tt.operators...)
			assert.Equal(t, tt.expected, got)
		})
	}
}
