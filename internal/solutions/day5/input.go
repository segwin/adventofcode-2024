package day5

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"strings"

	"github.com/segwin/adventofcode-2024/internal/transform"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	reachedUpdates := false
	sc := bufio.NewScanner(bytes.NewReader(inputData))
	for sc.Scan() {
		switch line := sc.Text(); {
		case line == "":
			reachedUpdates = true // empty line signals transition from rules to updates

		case reachedUpdates: // now parsing updates
			update, err := transform.Atois(strings.Split(line, ",")...)
			if err != nil {
				return nil, fmt.Errorf("parsing update line as ints (%q): %w", line, err)
			}
			s.Updates = append(s.Updates, update)

		default: // still parsing rules
			rules, err := transform.Atois(strings.SplitN(line, "|", 2)...)
			if err != nil {
				return nil, fmt.Errorf("parsing rule line as ints (%q): %w", line, err)
			}
			s.Rules = append(s.Rules, PageRule{Before: rules[0], After: rules[1]})
		}
	}

	return &s, nil
}
