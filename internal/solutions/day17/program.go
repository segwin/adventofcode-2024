package day17

import (
	"fmt"
	"strconv"
)

// ProgramState is the state of a program executed by the computer.
type ProgramState struct {
	// RegisterA, RegisterB and RegisterC are the computer's current register values.
	RegisterA, RegisterB, RegisterC int

	// Program to be executed.
	Program []Instruction

	// instructionPtr is the index of the next instruction to run.
	instructionPtr int
}

// Next instruction to be executed. If the program has completed, ok is set to false.
func (s *ProgramState) Next() (op Operation, operand Instruction, ok bool) {
	if s.instructionPtr < 0 || s.instructionPtr >= len(s.Program) {
		return 0, 0, false
	}
	return Operation(s.Program[s.instructionPtr]), s.Program[s.instructionPtr+1], true
}

// Instruction for execution by the computer, given as a 3-bit integer (0-7).
type Instruction int8

// Operation is a code defining the behaviour to execute.
type Operation Instruction

const (
	// ADV divides the value in register A divided by 2^(combo operand), truncated to an integer.
	// The result is stored in register A.
	ADV Operation = 0

	// BXL computes the bitwise XOR of register B and its literal operand.
	// The result is stored in register B.
	BXL Operation = 1

	// BST calculates (combo operand) % 8.
	// The result is stored in register B.
	BST Operation = 2

	// JNZ:
	//  - if register A == 0: does nothing
	//  - if register A != 0: jumps (sets the instruction pointer) to its literal operand
	//
	// If a jump occurs, the instruction pointer *is not* incremented by 2 following the operation.
	JNZ Operation = 3

	// BXL computes the bitwise XOR of register B and register C. Its operand is read but unused.
	// The result is stored in register B.
	BXC Operation = 4

	// OUT outputs (combo operand) % 8. If multiple values are outputted, they are separated by commas.
	OUT Operation = 5

	// BDV works like ADV but stores its result in register B. The numerator is still register A.
	BDV Operation = 6

	// CDV works like ADV but stores its result in register C. The numerator is still register A.
	CDV Operation = 7
)

// Execute this operation, returning the new program state.
func (c Operation) Execute(operand Instruction, cur ProgramState) (next ProgramState, output string) {
	next = cur
	next.instructionPtr += 2

	switch c {
	case ADV:
		next.RegisterA = cur.RegisterA / (1 << resolveComboOperand(operand, cur))

	case BXL:
		next.RegisterB = cur.RegisterB ^ int(operand)

	case BST:
		next.RegisterB = resolveComboOperand(operand, cur) % 8

	case JNZ:
		if cur.RegisterA != 0 {
			next.instructionPtr = int(operand)
		}

	case BXC:
		next.RegisterB = cur.RegisterB ^ cur.RegisterC

	case OUT:
		output = strconv.Itoa(resolveComboOperand(operand, cur) % 8)

	case BDV:
		next.RegisterB = cur.RegisterA / (1 << resolveComboOperand(operand, cur))

	case CDV:
		next.RegisterC = cur.RegisterA / (1 << resolveComboOperand(operand, cur))

	default:
		panic(fmt.Sprintf("invalid op code: %q", c))
	}

	return next, output
}

func resolveComboOperand(operand Instruction, state ProgramState) int {
	//   - Combo operands 0 through 3 represent literal values 0 through 3.
	//   - Combo operand 4 represents the value of register A.
	//   - Combo operand 5 represents the value of register B.
	//   - Combo operand 6 represents the value of register C.
	//   - Combo operand 7 is reserved and will not appear in valid programs.
	switch operand {
	case 0, 1, 2, 3:
		return int(operand) // literal value
	case 4:
		return state.RegisterA
	case 5:
		return state.RegisterB
	case 6:
		return state.RegisterC
	default: // 7
		panic("invalid combo operand: got reserved operand 7")
	}
}
