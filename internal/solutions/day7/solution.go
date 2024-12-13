package day7

import "fmt"

type Solution struct {
	OperandsByResult map[int][]int
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 7:\n")

	totalCalibration := TotalCalibration(s.OperandsByResult)
	fmt.Print("  PART 1:\n")
	fmt.Printf("    Total calibration result: %d\n", totalCalibration)

	return nil
}

func TotalCalibration(operandsByResult map[int][]int) int {
	sum := 0
	for result, operands := range operandsByResult {
		if !canBeTrue(result, operands) {
			continue
		}
		sum += result // ok: add test value to overall sum
	}
	return sum
}

func canBeTrue(expectedResult int, operands []int) bool {
	for _, operators := range OperatorCombinations(len(operands)) {
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
