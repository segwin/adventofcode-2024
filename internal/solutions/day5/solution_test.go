package day5_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSumCorrectUpdates(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		updates [][]int
		rules   []day5.PageRule

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example": {
			updates: [][]int{
				{75, 47, 61, 53, 29},
				{97, 61, 53, 29, 13},
				{75, 29, 13},
				{75, 97, 47, 61, 53},
				{61, 13, 29},
				{97, 13, 75, 29, 47},
			},
			rules: []day5.PageRule{
				{Before: 47, After: 53},
				{Before: 97, After: 13},
				{Before: 97, After: 61},
				{Before: 97, After: 47},
				{Before: 75, After: 29},
				{Before: 61, After: 13},
				{Before: 75, After: 53},
				{Before: 29, After: 13},
				{Before: 97, After: 29},
				{Before: 53, After: 29},
				{Before: 61, After: 53},
				{Before: 97, After: 53},
				{Before: 61, After: 29},
				{Before: 47, After: 13},
				{Before: 75, After: 47},
				{Before: 97, After: 75},
				{Before: 47, After: 61},
				{Before: 75, After: 61},
				{Before: 47, After: 29},
				{Before: 75, After: 13},
				{Before: 53, After: 13},
			},
			expected: 143,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day5.SumCorrectUpdates(tt.updates, tt.rules)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
