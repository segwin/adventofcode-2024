package day3_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day3.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestSumMuls(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		memory           string
		withConditionals bool

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example": {
			memory:           `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`,
			withConditionals: false,
			expected:         161,
		},
		"ok: day's example, with conditionals": {
			memory:           `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`,
			withConditionals: true,
			expected:         48,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day3.SumMuls(tt.memory, tt.withConditionals)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
