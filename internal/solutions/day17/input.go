package day17

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	for sc.Scan() {
		if err := parseLine(sc.Text(), &s.InitialState); err != nil {
			return nil, err
		}
	}

	return &s, nil
}

var (
	registerPattern = regexp.MustCompile(`Register ([ABC]): ([0-9]+)`)
	programPattern  = regexp.MustCompile(`Program: ([0-9](?:,[0-9])*)`)
)

func parseLine(line string, dst *ProgramState) (err error) {
	if matches := registerPattern.FindStringSubmatch(line); matches != nil {
		val, err := strconv.Atoi(matches[2])
		if err != nil {
			return fmt.Errorf("parsing register %s value %q: %w", matches[1], matches[2], err)
		}

		switch matches[1] {
		case "A":
			dst.RegisterA = val
		case "B":
			dst.RegisterB = val
		default: // "C"
			dst.RegisterC = val
		}
		return nil
	}

	if matches := programPattern.FindStringSubmatch(line); matches != nil {
		instructions, err := transform.Atois(strings.Split(matches[1], ",")...)
		if err != nil {
			return fmt.Errorf("parsing instructions %q: %w", matches[1], err)
		}

		dst.Program = make([]Instruction, len(instructions))
		for i, v := range instructions {
			if v < 0 || v > 7 {
				return fmt.Errorf("%w: instruction %v is out of range", parsing.ErrInvalidData, v)
			}
			dst.Program[i] = Instruction(v)
		}
	}

	return nil // ignore other lines, including newlines
}
