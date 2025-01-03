package day17_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day17"
	"github.com/stretchr/testify/assert"
)

func TestOperation_Execute(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		operation day17.Operation
		operand   day17.Instruction
		cur       day17.State

		// outputs
		expectedState  day17.State
		expectedOutput *day17.Instruction
	}{
		"day's instruction example 1": {
			operation: 2, operand: 6,
			cur: day17.State{RegisterA: 0, RegisterB: 0, RegisterC: 9},
			expectedState: day17.State{
				RegisterA: 0, RegisterB: 1, RegisterC: 9,
			},
			expectedOutput: nil,
		},
		"day's instruction example 4": {
			operation: 1, operand: 7,
			cur:            day17.State{RegisterA: 0, RegisterB: 29, RegisterC: 0},
			expectedState:  day17.State{RegisterA: 0, RegisterB: 26, RegisterC: 0},
			expectedOutput: nil,
		},
		"day's instruction example 5": {
			operation: 4, operand: 0,
			cur:            day17.State{RegisterA: 0, RegisterB: 2024, RegisterC: 43690},
			expectedState:  day17.State{RegisterA: 0, RegisterB: 44354, RegisterC: 43690},
			expectedOutput: nil,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotState, gotOutput := tt.operation.Execute(tt.operand, tt.cur)
			assert.Empty(t, cmp.Diff(tt.expectedState, gotState, cmpopts.IgnoreUnexported(day17.State{})))
			assert.Equal(t, tt.expectedOutput, gotOutput)
		})
	}
}
