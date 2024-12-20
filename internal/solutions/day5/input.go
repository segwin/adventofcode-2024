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
	s := Solution{
		PagesAfter: map[int][]int{},
		Updates:    nil,
	}

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

			page, after := rules[0], rules[1]
			s.PagesAfter[page] = append(s.PagesAfter[page], after)
		}
	}

	return &s, nil
}
