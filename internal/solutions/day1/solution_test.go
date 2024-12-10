package day1_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTotalDistance(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		left  []int
		right []int

		// outputs
		expected    int
		expectedErr error
	}{
		"error: mismatched left/right lengths": {
			left:        []int{1, 2, 3},
			right:       []int{4, 5},
			expectedErr: day1.ErrMismatchedLens,
		},
		"ok: day's example": {
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 11,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day1.TotalDistance(tt.left, tt.right)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestSimilarity(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		left  []int
		right []int

		// outputs
		expected    int
		expectedErr error
	}{
		"error: mismatched left/right lengths": {
			left:        []int{1, 2, 3},
			right:       []int{4, 5},
			expectedErr: day1.ErrMismatchedLens,
		},
		"ok: day's example": {
			left:     []int{3, 4, 2, 1, 3, 3},
			right:    []int{4, 3, 5, 3, 9, 3},
			expected: 31,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day1.Similarity(tt.left, tt.right)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
