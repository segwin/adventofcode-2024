package day1_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution_TotalDistance(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		solution day1.Solution

		// outputs
		expected    int
		expectedErr error
	}{
		"error: mismatched left/right lengths": {
			solution: day1.Solution{
				Left:  []int{1, 2, 3},
				Right: []int{4, 5},
			},
			expectedErr: day1.ErrMismatchedLens,
		},
		"ok: day's example": {
			solution: day1.Solution{
				Left:  []int{3, 4, 2, 1, 3, 3},
				Right: []int{4, 3, 5, 3, 9, 3},
			},
			expected: 11,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.solution.TotalDistance()
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSolution_Similarity(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		solution day1.Solution

		// outputs
		expected    int
		expectedErr error
	}{
		"error: mismatched left/right lengths": {
			solution: day1.Solution{
				Left:  []int{1, 2, 3},
				Right: []int{4, 5},
			},
			expectedErr: day1.ErrMismatchedLens,
		},
		"ok: day's example": {
			solution: day1.Solution{
				Left:  []int{3, 4, 2, 1, 3, 3},
				Right: []int{4, 3, 5, 3, 9, 3},
			},
			expected: 31,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := tt.solution.Similarity()
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
