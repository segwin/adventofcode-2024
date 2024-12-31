package day9

import (
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"io"
	"strconv"
)

//go:embed input.txt
var inputData []byte

func BuildSolution() (*Solution, error) {
	var s Solution

	r := bytes.NewReader(inputData)
	for {
		b, err := r.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, fmt.Errorf("reading input data: %w", err)
		}

		val, err := strconv.Atoi(string(b))
		if err != nil {
			return nil, fmt.Errorf("invalid input data: %w", err)
		}

		s.DiskMap = append(s.DiskMap, val)
	}

	return &s, nil
}
