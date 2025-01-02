package day15

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"

	"github.com/segwin/adventofcode-2024/internal/parsing"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	sc := bufio.NewScanner(bytes.NewReader(inputData))
	reachedInstructions := false
	for sc.Scan() {
		line := sc.Bytes()
		switch {
		case len(line) == 0:
			reachedInstructions = true

		case reachedInstructions:
			dirs, err := parseInstructions(line)
			if err != nil {
				return nil, err
			}
			s.Instructions = append(s.Instructions, dirs...)

		default: // still building layout
			row := make([]Tile, len(line))
			for i, t := range line {
				row[i] = Tile(t)
			}
			s.Layout.Map = append(s.Layout.Map, row)
		}
	}

	return &s, nil
}

func parseInstruction(instruction byte) (map2d.Direction, error) {
	switch instruction {
	case '^':
		return map2d.North(), nil
	case '>':
		return map2d.East(), nil
	case 'v':
		return map2d.South(), nil
	case '<':
		return map2d.West(), nil
	default:
		return map2d.Direction{}, fmt.Errorf("%w: unknown instruction %q", parsing.ErrInvalidData, instruction)
	}
}

func parseInstructions(instructions []byte) (out []map2d.Direction, err error) {
	out = make([]map2d.Direction, len(instructions))
	for i, instruction := range instructions {
		out[i], err = parseInstruction(instruction)
		if err != nil {
			return nil, fmt.Errorf("parsing instruction at offset %d: %w", i, err)
		}
	}
	return out, nil
}
