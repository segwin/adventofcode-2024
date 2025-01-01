package day5_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day5.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestSumCorrectUpdates(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		updates [][]int
		rules   map[int][]int

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
			rules: map[int][]int{
				47: {53, 13, 61, 29},
				97: {13, 61, 47, 29, 53, 75},
				75: {29, 53, 47, 61, 13},
				61: {13, 53, 29},
				29: {13},
				53: {29, 13},
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

func TestSumIncorrectUpdates(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		updates [][]int
		rules   map[int][]int

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
			rules: map[int][]int{
				47: {53, 13, 61, 29},
				97: {13, 61, 47, 29, 53, 75},
				75: {29, 53, 47, 61, 13},
				61: {13, 53, 29},
				29: {13},
				53: {29, 13},
			},
			expected: 123,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day5.SumIncorrectUpdates(tt.updates, tt.rules)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
