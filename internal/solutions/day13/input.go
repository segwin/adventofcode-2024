package day13

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	var cur *ClawMachine
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			// machine description is complete
			s.Machines = append(s.Machines, *cur)
			cur = nil
			continue
		}

		if cur == nil {
			cur = &ClawMachine{} // new machine is being described
		}

		if err := parseLine(line, cur); err != nil {
			return nil, fmt.Errorf("parsing line %q: %w", line, err)
		}
	}

	return &s, nil
}

func parseDistance(xstr, ystr string) (map2d.Distance, error) {
	x, err := strconv.Atoi(xstr)
	if err != nil {
		return map2d.Distance{}, fmt.Errorf("parsing x value %q: %w", xstr, err)
	}
	y, err := strconv.Atoi(ystr)
	if err != nil {
		return map2d.Distance{}, fmt.Errorf("parsing y value %q: %w", ystr, err)
	}
	return map2d.Distance{X: x, Y: y}, nil
}

var (
	moveAPattern = regexp.MustCompile(`Button A: X\+([0-9]+), Y\+([0-9]+)`)
	moveBPattern = regexp.MustCompile(`Button B: X\+([0-9]+), Y\+([0-9]+)`)
	prizePattern = regexp.MustCompile(`Prize: X=([0-9]+), Y=([0-9]+)`)
)

func parseLine(line string, dst *ClawMachine) (err error) {
	if matches := moveAPattern.FindStringSubmatch(line); matches != nil {
		dst.MoveA, err = parseDistance(matches[1], matches[2])
		if err != nil {
			return err
		}
		return nil
	}
	if matches := moveBPattern.FindStringSubmatch(line); matches != nil {
		dst.MoveB, err = parseDistance(matches[1], matches[2])
		if err != nil {
			return err
		}
		return nil
	}
	if matches := prizePattern.FindStringSubmatch(line); matches != nil {
		dst.Prize, err = parseDistance(matches[1], matches[2])
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%w: unknown line pattern", parsing.ErrInvalidData)
}
