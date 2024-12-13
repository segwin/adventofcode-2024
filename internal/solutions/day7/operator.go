package day7

import (
	"math"
)

type Operator interface {
	// Symbol for this operator for use in equations.
	Symbol() string
	// Apply this operator to the given left & right operands.
	Apply(l, r int) int
}

// Plus is the + operator. It adds two operands together.
type Plus struct{}

func (Plus) Symbol() string     { return "+" }
func (Plus) Apply(l, r int) int { return l + r }

// Star is the * operator. It multiplies two operands together.
type Star struct{}

func (Star) Symbol() string     { return "*" }
func (Star) Apply(l, r int) int { return l * r }

// DoublePipe is the || operator. It concatenates two operands together.
type DoublePipe struct{}

func (DoublePipe) Symbol() string { return "||" }
func (DoublePipe) Apply(l, r int) int {
	// avoid string conversion by using some simple math
	// note: this won't work for negative or decimal numbers, but we only need to handle positive integers
	if r == 0 {
		return l * 10 // special case: log10(0+1) == 0, but we may want to include 0 in our problem set
	}
	leftPow10 := int(math.Ceil(math.Log10(float64(r + 1))))
	return l*int(math.Pow10(leftPow10)) + r
}

// OperatorCombinations returns all possible sets of operators for the given number of operands.
func OperatorCombinations(numOperands int, operators []Operator) [][]Operator {
	if numOperands == 0 || numOperands == 1 {
		return nil // no operators in these cases
	}

	numOperators := numOperands - 1
	numCombinations := int(math.Round(math.Pow(float64(len(operators)), float64(numOperators))))
	combinations := make([][]Operator, numCombinations)
	for i := range combinations {
		nitset := toBaseN(len(operators), i, numOperators)

		combinations[i] = make([]Operator, numOperators)
		for j := range combinations[i] {
			combinations[i][j] = operators[nitset[j]%len(operators)]
		}
	}
	return combinations
}

// toBaseN returns the base N representation of val, given as a list of "nits" (e.g. bits for binary).
// The resulting array is padded with zeros to meet the given size.
func toBaseN(n int, val int, size int) []int {
	nits := make([]int, 0, size)
	for remaining := val; remaining != 0; remaining /= n {
		nits = append(nits, remaining%n)
	}

	// pad remainder with 0s
	for i := len(nits); i < size; i++ {
		nits = append(nits, 0)
	}

	return nits
}
