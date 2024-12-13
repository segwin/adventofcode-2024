package day7

import (
	"fmt"
	"sync"
)

type Solution struct {
	OperandsByResult map[int][]int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 7:\n")

	fmt.Print("  PART 1:\n")
	fmt.Printf("    +, *: %d\n", TotalCalibration(s.OperandsByResult, Plus{}, Star{}))

	fmt.Print("  PART 2:\n")
	fmt.Printf("    +, *, ||: %d\n", TotalCalibration(s.OperandsByResult, Plus{}, Star{}, DoublePipe{}))

	return nil
}

func TotalCalibration(operandsByResult map[int][]int, operators ...Operator) int {
	// check all combinations in parallel
	var wg sync.WaitGroup
	calibratedValues := make(chan int)
	for result, operands := range operandsByResult {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if canBeTrue(result, operands, operators) {
				calibratedValues <- result // ok: add test value to overall sum
			}
		}()
	}

	// once all producers are done, close the channel to end collection
	go func() {
		wg.Wait()
		close(calibratedValues)
	}()

	// collect results
	sum := 0
	for val := range calibratedValues {
		sum += val
	}

	return sum
}

func canBeTrue(expectedResult int, operands []int, operators []Operator) bool {
	for _, operators := range OperatorCombinations(len(operands), operators) {
		result := operands[0]
		for i := 0; i < len(operators); i++ { // evaluate operators left-to-right
			result = operators[i].Apply(result, operands[i+1])
		}

		// fully evaluated: check if value meets expected result
		if result == expectedResult {
			return true
		}
	}
	return false
}
