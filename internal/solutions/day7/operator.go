package day7

import "math/big"

type BinaryOperator interface {
	// Apply this operator to the given left & right operands.
	Apply(l, r int) int
}

type Plus struct{}

func (Plus) Apply(l, r int) int { return l + r }

type Times struct{}

func (Times) Apply(l, r int) int { return l * r }

// OperatorCombinations returns all possible sets of operators for the given number of operands.
func OperatorCombinations(numOperands int) [][]BinaryOperator {
	if numOperands == 0 || numOperands == 1 {
		return nil // no operators in these cases
	}

	numOperators := numOperands - 1
	numCombinations := 1 << numOperators // 2^(numOperators)
	combinations := make([][]BinaryOperator, numCombinations)
	for i := range combinations {
		bitset := big.NewInt(int64(i))

		combinations[i] = make([]BinaryOperator, numOperators)
		for j := range combinations[i] {
			switch bitset.Bit(len(combinations[i]) - 1 - j) { // cosmetic: use LSB order to please my OCD
			case 0:
				combinations[i][j] = Plus{}
			case 1:
				combinations[i][j] = Times{}
			}
		}
	}
	return combinations
}
