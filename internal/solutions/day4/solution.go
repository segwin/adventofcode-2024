package day4

import (
	"bytes"
	"fmt"
)

type Solution struct {
	Search []string
}

func (s *Solution) RunToConsole() error {
	fmt.Print("DAY 4:\n")

	count := CountXMAS(s.Search)

	fmt.Print("  PART 1:\n")
	fmt.Printf("    Count: %d\n", count)
	return nil
}

func CountXMAS(search []string) int {
	count := 0
	for i := range search {
		for j := range search[i] {
			c1 := search[i][j]
			if c1 != 'X' { // only start searching around X's
				continue
			}

			count += xmasAround(search, i, j)
		}
	}
	return count
}

func isXMAS(b1, b2, b3, b4 byte) bool {
	return bytes.Equal([]byte("XMAS"), []byte{b1, b2, b3, b4})
}

func xmasAround(search []string, i, j int) int {
	count := 0

	if j >= 3 && isXMAS(search[i][j], search[i][j-1], search[i][j-2], search[i][j-3]) {
		count++ // ←
	}
	if i >= 3 && j >= 3 && isXMAS(search[i][j], search[i-1][j-1], search[i-2][j-2], search[i-3][j-3]) {
		count++ // ↖
	}
	if i >= 3 && isXMAS(search[i][j], search[i-1][j], search[i-2][j], search[i-3][j]) {
		count++ // ↑
	}
	if i >= 3 && j <= len(search)-4 && isXMAS(search[i][j], search[i-1][j+1], search[i-2][j+2], search[i-3][j+3]) {
		count++ // ↗
	}
	if j <= len(search[i])-4 && isXMAS(search[i][j], search[i][j+1], search[i][j+2], search[i][j+3]) {
		count++ // →
	}
	if i <= len(search)-4 && j <= len(search[i])-4 && isXMAS(search[i][j], search[i+1][j+1], search[i+2][j+2], search[i+3][j+3]) {
		count++ // ↘
	}
	if i <= len(search)-4 && isXMAS(search[i][j], search[i+1][j], search[i+2][j], search[i+3][j]) {
		count++ // ↓
	}
	if i <= len(search)-4 && j >= 3 && isXMAS(search[i][j], search[i+1][j-1], search[i+2][j-2], search[i+3][j-3]) {
		count++ // ↙
	}

	return count
}
