package day17

import (
	"fmt"
	"strings"
)

type Solution struct {
	InitialState ProgramState
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 17:\n")

	_, output := ExecuteProgram(s.InitialState)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Program output: %s\n", output)

	fmt.Print("  PART 2:\n")
	fmt.Printf("    TODO\n")

	return nil
}

// ExecuteProgram runs the given program through to completion and returns the resulting output.
func ExecuteProgram(state ProgramState) (finalState ProgramState, output string) {
	var outputs []string

	for {
		op, operand, ok := state.Next()
		if !ok {
			return state, strings.Join(outputs, ",") // done: return constructed output
		}

		var opOutput string
		state, opOutput = op.Execute(operand, state)
		if opOutput != "" {
			outputs = append(outputs, opOutput)
		}
	}
}
