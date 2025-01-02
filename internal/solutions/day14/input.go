package day14

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/parsing"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	r := parsing.SSVReader{Data: bytes.NewReader(inputData), ExpectedCols: 2}
	for row, err := range r.All() {
		if err != nil {
			return nil, fmt.Errorf("parsing input data: %w", err)
		}

		var robot RobotState
		robot.Position.X, robot.Position.Y, err = parseXY(row[0], "p=")
		if err != nil {
			return nil, fmt.Errorf("parsing entry %q: %w", row[0], err)
		}
		robot.Velocity.X, robot.Velocity.Y, err = parseXY(row[1], "v=")
		if err != nil {
			return nil, fmt.Errorf("parsing entry %q: %w", row[1], err)
		}

		s.Robots = append(s.Robots, robot)
	}

	return &s, nil
}

func parseXY(entry string, expectedPrefix string) (x, y int, err error) {
	csv := strings.TrimPrefix(entry, expectedPrefix)
	if csv == entry {
		return 0, 0, fmt.Errorf("%w: missing expected prefix %q", parsing.ErrInvalidData, expectedPrefix)
	}

	xstr, ystr, ok := strings.Cut(csv, ",")
	if !ok {
		return 0, 0, fmt.Errorf("%w: does not contain exactly 2 comma-separated values", parsing.ErrInvalidData)
	}

	x, err = strconv.Atoi(xstr)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing x value %q: %w", xstr, err)
	}
	y, err = strconv.Atoi(ystr)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing x value %q: %w", ystr, err)
	}
	return x, y, nil
}
