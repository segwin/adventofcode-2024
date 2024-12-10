package day2_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution_SafeReports(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		solution         day2.Solution
		problemDampening bool

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example, no problem dampening": {
			solution: day2.Solution{
				Reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			problemDampening: false,
			expected:         2,
		},
		"ok: day's example, with problem dampening": {
			solution: day2.Solution{
				Reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
			},
			problemDampening: true,
			expected:         4,
		},
		"ok: problem dampening where removal is not on first unsafe index": {
			solution: day2.Solution{
				Reports: [][]int{
					{4, 5, 4, 7},
				},
			},
			problemDampening: true,
			expected:         1,
		},
		"ok: report with <2 elements are always safe": {
			solution: day2.Solution{
				Reports: [][]int{
					{},
					{123},
				},
			},
			problemDampening: true,
			expected:         2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.solution.SafeReports(tt.problemDampening)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
