package day17_test

import (
	"testing"

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
		initialState day17.State
		program      []day17.Instruction

		// outputs
		expected string
	}{
		"day's instruction example 2": {
			program:      []day17.Instruction{5, 0, 5, 1, 5, 4},
			initialState: day17.State{RegisterA: 10, RegisterB: 0, RegisterC: 0},
			expected:     "0,1,2",
		},
		"day's instruction example 3": {
			initialState: day17.State{RegisterA: 2024, RegisterB: 0, RegisterC: 0},
			program:      []day17.Instruction{0, 1, 5, 4, 3, 0},
			expected:     "4,2,5,6,7,7,7,7,3,1,0",
		},

		"day's full example": {
			initialState: day17.State{RegisterA: 729, RegisterB: 0, RegisterC: 0},
			program:      []day17.Instruction{0, 1, 5, 4, 3, 0},
			expected:     "4,6,3,5,6,3,5,2,1,0",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day17.ExecuteProgram(tt.program, tt.initialState)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestCorrectRegisterA(t *testing.T) {
	t.Parallel()

	program := []day17.Instruction{0, 3, 5, 4, 3, 0}
	state := day17.State{RegisterA: 2024, RegisterB: 0, RegisterC: 0}

	got := day17.CorrectRegisterA(program, state)
	assert.Equal(t, 117440, got)
}
