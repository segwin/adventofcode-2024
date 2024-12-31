package day11_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day11"
	"github.com/stretchr/testify/assert"
)

func TestBlink(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		stones []day11.Stone

		// outputs
		expected []day11.Stone
	}{
		"day's simplified example": {
			stones:   []day11.Stone{0, 1, 10, 99, 999},
			expected: []day11.Stone{1, 2024, 1, 0, 9, 9, 2021976},
		},
		"day's full example: blink 1": {
			stones:   []day11.Stone{125, 17},
			expected: []day11.Stone{253000, 1, 7},
		},
		"day's full example: blink 2": {
			stones:   []day11.Stone{253000, 1, 7},
			expected: []day11.Stone{253, 0, 2024, 14168},
		},
		"day's full example: blink 3": {
			stones:   []day11.Stone{253, 0, 2024, 14168},
			expected: []day11.Stone{512072, 1, 20, 24, 28676032},
		},
		"day's full example: blink 4": {
			stones:   []day11.Stone{512072, 1, 20, 24, 28676032},
			expected: []day11.Stone{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
		},
		"day's full example: blink 5": {
			stones:   []day11.Stone{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
			expected: []day11.Stone{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
		},
		"day's full example: blink 6": {
			stones:   []day11.Stone{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
			expected: []day11.Stone{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day11.Blink(tt.stones)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestBlinkTimes(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		stones []day11.Stone
		times  int

		// outputs
		expectedCount int
	}{
		"day's example: 6 blinks": {
			stones:        []day11.Stone{125, 17},
			times:         6,
			expectedCount: 22,
		},
		"day's example: 25 blinks": {
			stones:        []day11.Stone{125, 17},
			times:         25,
			expectedCount: 55312,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day11.BlinkTimes(tt.stones, tt.times)
			assert.Equal(t, tt.expectedCount, len(got))
		})
	}
}
