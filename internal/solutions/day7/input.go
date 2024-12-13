package day7

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	s := Solution{
		OperandsByResult: map[int][]int{},
	}

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	for sc.Scan() {
		line := sc.Text()
		resultStr, operandsSSV, _ := strings.Cut(line, ":")

		result, err := strconv.Atoi(resultStr)
		if err != nil {
			return nil, fmt.Errorf("parsing result on line %q: %w", line, err)
		}

		operands, err := transform.Atois(strings.Fields(strings.TrimSpace(operandsSSV))...)
		if err != nil {
			return nil, fmt.Errorf("parsing operands on line %q: %w", line, err)
		}

		s.OperandsByResult[result] = operands
	}

	return &s, nil
}
