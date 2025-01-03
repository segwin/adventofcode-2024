package day17

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var (
	errLoopDetected = errors.New("program entered a loop and was terminated")
	errPanicked     = errors.New("program was terminated due to a panic")
	errEarlyExit    = errors.New("program was terminated early to to an exit condition")
)

type Solution struct {
	InitialState State
	Program      []Instruction
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 17:\n")

	output, err := ExecuteProgram(s.Program, s.InitialState)
	if err != nil {
		return fmt.Errorf("program failed: %w", err)
	}

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Program output: %s\n", output)

	fmt.Print("  PART 2:\n")
	fmt.Printf("    Correct register A value: %d\n", CorrectRegisterA(s.Program, s.InitialState))

	return nil
}

// ExecuteProgram runs the given program through to completion and returns the resulting output.
// If a loop is detected or the program otherwise exits, ok is set to false.
func ExecuteProgram(program []Instruction, state State) (output string, err error) {
	rawOutput, err := executeProgram(program, state, map[State]State{}, nil)
	if err != nil {
		return "", err
	}
	return instructionsToCSV(rawOutput), nil
}

// CorrectRegisterA returns the lowest positive register A value that causes the given program
// to output itself.
func CorrectRegisterA(program []Instruction, state State) int {
	// cache := map[State]State{}
	// for {
	// 	rawOutput, err := executeProgram(program, state, cache, nil)
	// 	if err == nil && slices.Equal(rawOutput, program) {
	// 		return state.RegisterA
	// 	}
	// 	state.RegisterA += 1e3 // try again with the next value
	// }

	// the naive implementation above is too slow, so although a general-purpose solution is nice
	// it's best to have something that works in human-scale time
	return correctRegisterAEmpirical(program, state)
}

// correctRegisterAEmpirical uses empirically-derived knowledge of the input program's behaviour:
// the farther right in the output, the more stable the values.
//
// We use this knowledge to first find a suitable ballpark (correct # instructions) and progressively
// use smaller increments as right-hand instructions stabilise on the correct values.
func correctRegisterAEmpirical(program []Instruction, state State) int {
	// first, look for a value that produces the right number of output digits to get into the right ballpark
	cache := map[State]State{}
	state.RegisterA = 1
	fmt.Println("    Finding ballpark...")
	for {
		rawOutput, err := executeProgram(program, state, cache, nil)
		if err == nil && len(rawOutput) == len(program) {
			break // ok: we're at the right number of digits, try looking around this value
		}
		state.RegisterA *= 2 // try again with increasingly larger values
	}

	fmt.Printf("    Starting from %d, narrowing onto correct value...\n", state.RegisterA)
	for {
		rawOutput, err := executeProgram(program, state, cache, nil)
		if err != nil {
			state.RegisterA++ // failed, try immediate successor
			continue
		}
		if slices.Equal(rawOutput, program) {
			return state.RegisterA // found it!
		}
		if len(rawOutput) > len(program) {
			state.RegisterA /= 2 // went too far, go back
			continue
		}

		// reduce scaling amount as we approach more correct digits (right-to-left)
		firstIncorrectRTL := 0
		for i := len(program) - 1; i >= 0; i-- {
			if rawOutput[i] != program[i] {
				firstIncorrectRTL = i
				break
			}
		}
		state.RegisterA += 1 << (2 * firstIncorrectRTL) // scale increment based on how many right-hand digits we have right
	}
}

// executeProgram executes the given program, but with additional controls to make looking for the
// correct register A value more efficient.
//
//   - cache allows avoiding recomputation when a previously-calculated state is encountered in
//     a new executeProgram call. It is mutated by each call.
//   - exitEarly allows execution to exit early if the predicate is met, returning errEarlyExit.
func executeProgram(
	program []Instruction, state State,
	cache map[State]State,
	exitEarly func(rawOutput []Instruction) bool,
) (rawOutput []Instruction, err error) {
	defer func() {
		if details := recover(); details != nil {
			err = fmt.Errorf("%w: %v", errPanicked, details)
		}
	}()

	previousStates := map[State]struct{}{} // for loop detection
	for {
		// detect loops & exit immediately
		if _, inALoop := previousStates[state]; inALoop {
			return nil, errLoopDetected
		}
		previousStates[state] = struct{}{}

		// need to compute the operation
		operation, operand, ok := state.Next(program)
		if !ok {
			return rawOutput, nil // done: return constructed output
		}

		if operation != OUT {
			// for non-output operations, check to see if we can use a cached state
			// don't cache OUTs: it would make us sensitive to initial state changes, defeating the purpose
			if next, ok := cache[state]; ok {
				state = next
				continue
			}
		}

		next, opOutput := operation.Execute(operand, state)
		if opOutput != nil {
			rawOutput = append(rawOutput, *opOutput)
		}

		if operation != OUT {
			cache[state] = next // cache the operation for next time
		}
		state = next

		// can we exit early?
		if exitEarly != nil && exitEarly(rawOutput) {
			return nil, errEarlyExit
		}
	}
}

func instructionsToCSV(instructions []Instruction) string {
	strs := make([]string, len(instructions))
	for i, v := range instructions {
		strs[i] = strconv.Itoa(int(v))
	}
	return strings.Join(strs, ",")
}
