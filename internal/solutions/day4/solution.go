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

	xmases := CountXMAS(s.Search)
	fmt.Print("  PART 1:\n")
	fmt.Printf("    XMASes: %d\n", xmases)

	crossMases := CountCrossMas(s.Search)
	fmt.Print("  PART 2:\n")
	fmt.Printf("    X-MASes: %d\n", crossMases)
	return nil
}

// CountXMAS returns the number of times "XMAS" appears in the search grid, in any direction.
func CountXMAS(search []string) int {
	count := 0
	for i := range search {
		for j := range search[i] {
			if search[i][j] != 'X' { // only search around X's (start)
				continue
			}
			count += xmasAround(search, i, j)
		}
	}
	return count
}

// CountCrossMas returns the number of times "MAS" appears in an X shape in the search grid.
func CountCrossMas(search []string) int {
	count := 0
	for i := range search {
		for j := range search[i] {
			if search[i][j] != 'A' {
				continue // only search around A's (middle)
			}
			if isMASCenter(search, i, j) {
				count++
			}
		}
	}
	return count
}

func matches(expected string, data ...byte) bool {
	return bytes.Equal([]byte(expected), data)
}

func xmasAround(s []string, i, j int) int {
	count := 0

	if j >= 3 && matches("XMAS", s[i][j], s[i][j-1], s[i][j-2], s[i][j-3]) {
		count++ // ←
	}
	if i >= 3 && j >= 3 && matches("XMAS", s[i][j], s[i-1][j-1], s[i-2][j-2], s[i-3][j-3]) {
		count++ // ↖
	}
	if i >= 3 && matches("XMAS", s[i][j], s[i-1][j], s[i-2][j], s[i-3][j]) {
		count++ // ↑
	}
	if i >= 3 && j <= len(s)-4 && matches("XMAS", s[i][j], s[i-1][j+1], s[i-2][j+2], s[i-3][j+3]) {
		count++ // ↗
	}
	if j <= len(s[i])-4 && matches("XMAS", s[i][j], s[i][j+1], s[i][j+2], s[i][j+3]) {
		count++ // →
	}
	if i <= len(s)-4 && j <= len(s[i])-4 && matches("XMAS", s[i][j], s[i+1][j+1], s[i+2][j+2], s[i+3][j+3]) {
		count++ // ↘
	}
	if i <= len(s)-4 && matches("XMAS", s[i][j], s[i+1][j], s[i+2][j], s[i+3][j]) {
		count++ // ↓
	}
	if i <= len(s)-4 && j >= 3 && matches("XMAS", s[i][j], s[i+1][j-1], s[i+2][j-2], s[i+3][j-3]) {
		count++ // ↙
	}

	return count
}

func isMASCenter(s []string, i, j int) bool {
	if i == 0 || i == len(s)-1 || j == 0 || j == len(s[i])-1 {
		return false // can't be the middle of a MAS if there's no room to all sides
	}

	southeast := []byte{s[i-1][j-1], s[i][j], s[i+1][j+1]} // ↘
	northwest := []byte{s[i+1][j+1], s[i][j], s[i-1][j-1]} // ↖
	if !matches("MAS", southeast...) && !matches("MAS", northwest...) {
		return false
	}

	southwest := []byte{s[i-1][j+1], s[i][j], s[i+1][j-1]} // ↗
	northeast := []byte{s[i+1][j-1], s[i][j], s[i-1][j+1]} // ↙
	if !matches("MAS", southwest...) && !matches("MAS", northeast...) {
		return false
	}

	return true
}
