package day7_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day7"
	"github.com/stretchr/testify/assert"
)

func TestOperatorCombinations(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		numOperands int

		// outputs
		expected [][]day7.BinaryOperator
	}{
		"0 => 0": {
			numOperands: 0,
			expected:    nil,
		},
		"1 => 0": {
			numOperands: 1,
			expected:    nil,
		},
		"2 => 2": {
			numOperands: 2,
			expected: [][]day7.BinaryOperator{
				{day7.Plus{}},
				{day7.Times{}},
			},
		},
		"3 => 4": {
			numOperands: 3,
			expected: [][]day7.BinaryOperator{
				{day7.Plus{}, day7.Plus{}},
				{day7.Plus{}, day7.Times{}},
				{day7.Times{}, day7.Plus{}},
				{day7.Times{}, day7.Times{}},
			},
		},
		"4 => 8": {
			numOperands: 4,
			expected: [][]day7.BinaryOperator{
				{day7.Plus{}, day7.Plus{}, day7.Plus{}},
				{day7.Plus{}, day7.Plus{}, day7.Times{}},
				{day7.Plus{}, day7.Times{}, day7.Plus{}},
				{day7.Plus{}, day7.Times{}, day7.Times{}},
				{day7.Times{}, day7.Plus{}, day7.Plus{}},
				{day7.Times{}, day7.Plus{}, day7.Times{}},
				{day7.Times{}, day7.Times{}, day7.Plus{}},
				{day7.Times{}, day7.Times{}, day7.Times{}},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day7.OperatorCombinations(tt.numOperands)
			assert.Empty(t, cmp.Diff(tt.expected, got, cmpopts.EquateEmpty()))
		})
	}
}
