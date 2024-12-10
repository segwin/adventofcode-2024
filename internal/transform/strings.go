package transform

import (
	"fmt"
	"strconv"
)

// Atois applies strconv.Atoi to all values in the given list, stopping if any value is invalid.
func Atois(in ...string) ([]int, error) {
	out := make([]int, len(in))
	for i, str := range in {
		val, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("parsing column %d as int: %w", i, err)
		}
		out[i] = val
	}
	return out, nil
}
