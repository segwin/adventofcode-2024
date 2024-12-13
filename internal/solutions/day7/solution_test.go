package day7_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day7"
	"github.com/stretchr/testify/assert"
)

func TestTotalCalibration(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		operandsByResult map[int][]int

		// outputs
		expected int
	}{
		"ok: day's example": {
			operandsByResult: map[int][]int{
				190:    {10, 19},
				3267:   {81, 40, 27},
				83:     {17, 5},
				156:    {15, 6},
				7290:   {6, 8, 6, 15},
				161011: {16, 10, 13},
				192:    {17, 8, 14},
				21037:  {9, 7, 18, 13},
				292:    {11, 6, 16, 20},
			},
			expected: 3749,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day7.TotalCalibration(tt.operandsByResult)
			assert.Equal(t, tt.expected, got)
		})
	}
}
