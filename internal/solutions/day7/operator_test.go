package day7_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day7"
	"github.com/stretchr/testify/assert"
)

func TestOperators(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		op   day7.Operator
		l, r int

		// outputs
		expected int
	}{
		"0 + 0 => 0": {op: day7.Plus{}, l: 0, r: 0, expected: 0},
		"1 + 0 => 1": {op: day7.Plus{}, l: 1, r: 0, expected: 1},
		"2 + 5 => 7": {op: day7.Plus{}, l: 2, r: 5, expected: 7},

		"0 * 0 => 0":  {op: day7.Star{}, l: 0, r: 0, expected: 0},
		"1 * 0 => 0":  {op: day7.Star{}, l: 1, r: 0, expected: 0},
		"2 * 5 => 10": {op: day7.Star{}, l: 2, r: 5, expected: 10},

		"0 || 0 => 0":  {op: day7.DoublePipe{}, l: 0, r: 0, expected: 0},
		"1 || 0 => 10": {op: day7.DoublePipe{}, l: 1, r: 0, expected: 10},
		"2 || 5 => 25": {op: day7.DoublePipe{}, l: 2, r: 5, expected: 25},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.op.Apply(tt.l, tt.r)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestOperatorCombinations(t *testing.T) {
	t.Parallel()

	operators := []day7.Operator{day7.Plus{}, day7.Star{}, day7.DoublePipe{}}

	tests := map[string]struct {
		// inputs
		numOperands int

		// outputs
		expected [][]day7.Operator
	}{
		"0 => 0": {
			numOperands: 0,
			expected:    nil,
		},
		"1 => 0": {
			numOperands: 1,
			expected:    nil,
		},
		"2 => 3": {
			numOperands: 2,
			expected: [][]day7.Operator{
				{day7.Plus{}},
				{day7.Star{}},
				{day7.DoublePipe{}},
			},
		},
		"3 => 9": {
			numOperands: 3,
			expected: [][]day7.Operator{
				{day7.Plus{}, day7.Plus{}},
				{day7.Star{}, day7.Plus{}},
				{day7.DoublePipe{}, day7.Plus{}},
				{day7.Plus{}, day7.Star{}},
				{day7.Star{}, day7.Star{}},
				{day7.DoublePipe{}, day7.Star{}},
				{day7.Plus{}, day7.DoublePipe{}},
				{day7.Star{}, day7.DoublePipe{}},
				{day7.DoublePipe{}, day7.DoublePipe{}},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day7.OperatorCombinations(tt.numOperands, operators)
			assert.Empty(t, cmp.Diff(tt.expected, got, cmpopts.EquateEmpty()))
		})
	}
}
