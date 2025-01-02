package day17_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/segwin/adventofcode-2024/internal/solutions/day17"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day17.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestExecuteProgram(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		initialState day17.ProgramState

		// outputs
		expectedState  day17.ProgramState
		expectedOutput string
	}{
		"day's instruction example 1": {
			initialState: day17.ProgramState{
				RegisterA: 0, RegisterB: 0, RegisterC: 9,
				Program: []day17.Instruction{2, 6},
			},
			expectedState: day17.ProgramState{
				RegisterA: 0, RegisterB: 1, RegisterC: 9,
				Program: []day17.Instruction{2, 6},
			},
			expectedOutput: "",
		},
		"day's instruction example 2": {
			initialState: day17.ProgramState{
				RegisterA: 10, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{5, 0, 5, 1, 5, 4},
			},
			expectedState: day17.ProgramState{
				RegisterA: 10, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{5, 0, 5, 1, 5, 4},
			},
			expectedOutput: "0,1,2",
		},
		"day's instruction example 3": {
			initialState: day17.ProgramState{
				RegisterA: 2024, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{0, 1, 5, 4, 3, 0},
			},
			expectedState: day17.ProgramState{
				RegisterA: 0, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{0, 1, 5, 4, 3, 0},
			},
			expectedOutput: "4,2,5,6,7,7,7,7,3,1,0",
		},
		"day's instruction example 4": {
			initialState: day17.ProgramState{
				RegisterA: 0, RegisterB: 29, RegisterC: 0,
				Program: []day17.Instruction{1, 7},
			},
			expectedState: day17.ProgramState{
				RegisterA: 0, RegisterB: 26, RegisterC: 0,
				Program: []day17.Instruction{1, 7},
			},
			expectedOutput: "",
		},
		"day's instruction example 5": {
			initialState: day17.ProgramState{
				RegisterA: 0, RegisterB: 2024, RegisterC: 43690,
				Program: []day17.Instruction{4, 0},
			},
			expectedState: day17.ProgramState{
				RegisterA: 0, RegisterB: 44354, RegisterC: 43690,
				Program: []day17.Instruction{4, 0},
			},
			expectedOutput: "",
		},

		"day's full example": {
			initialState: day17.ProgramState{
				RegisterA: 729, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{0, 1, 5, 4, 3, 0},
			},
			expectedState: day17.ProgramState{
				RegisterA: 0, RegisterB: 0, RegisterC: 0,
				Program: []day17.Instruction{0, 1, 5, 4, 3, 0},
			},
			expectedOutput: "4,6,3,5,6,3,5,2,1,0",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			gotState, gotOutput := day17.ExecuteProgram(tt.initialState)
			assert.Empty(t, cmp.Diff(tt.expectedState, gotState, cmpopts.IgnoreUnexported(day17.ProgramState{})))
			assert.Empty(t, cmp.Diff(tt.expectedOutput, gotOutput))
		})
	}
}
