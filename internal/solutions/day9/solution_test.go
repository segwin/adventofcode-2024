package day9_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/day9"
	"github.com/stretchr/testify/assert"
)

func TestParseDiskMap(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		diskMap []int

		// outputs
		expected []*int
	}{
		"day's simplified example": {
			diskMap: []int{1, 2, 3, 4, 5},
			expected: []*int{
				ptr(0),
				nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil, nil,
				ptr(2), ptr(2), ptr(2), ptr(2), ptr(2),
			},
		},
		"day's full example": {
			diskMap: []int{2, 3, 3, 3, 1, 3, 3, 1, 2, 1, 4, 1, 4, 1, 3, 1, 4, 0, 2},
			expected: []*int{
				ptr(0), ptr(0),
				nil, nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil,
				ptr(2),
				nil, nil, nil,
				ptr(3), ptr(3), ptr(3),
				nil,
				ptr(4), ptr(4),
				nil,
				ptr(5), ptr(5), ptr(5), ptr(5),
				nil,
				ptr(6), ptr(6), ptr(6), ptr(6),
				nil,
				ptr(7), ptr(7), ptr(7),
				nil,
				ptr(8), ptr(8), ptr(8), ptr(8),
				ptr(9), ptr(9),
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day9.ParseDiskMap(tt.diskMap)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestCompactBlocks(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		uncompacted []*int

		// outputs
		expected []*int
	}{
		"day's simplified example": {
			uncompacted: []*int{
				ptr(0),
				nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil, nil,
				ptr(2), ptr(2), ptr(2), ptr(2), ptr(2),
			},
			expected: []*int{
				ptr(0),
				ptr(2), ptr(2),
				ptr(1), ptr(1), ptr(1),
				ptr(2), ptr(2), ptr(2), nil,
				nil, nil, nil, nil, nil,
			},
		},
		"day's full example": {
			uncompacted: []*int{
				ptr(0), ptr(0),
				nil, nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil,
				ptr(2),
				nil, nil, nil,
				ptr(3), ptr(3), ptr(3),
				nil,
				ptr(4), ptr(4),
				nil,
				ptr(5), ptr(5), ptr(5), ptr(5),
				nil,
				ptr(6), ptr(6), ptr(6), ptr(6),
				nil,
				ptr(7), ptr(7), ptr(7),
				nil,
				ptr(8), ptr(8), ptr(8), ptr(8),
				ptr(9), ptr(9),
			},
			expected: []*int{
				ptr(0), ptr(0),
				ptr(9), ptr(9), ptr(8),
				ptr(1), ptr(1), ptr(1),
				ptr(8), ptr(8), ptr(8),
				ptr(2),
				ptr(7), ptr(7), ptr(7),
				ptr(3), ptr(3), ptr(3),
				ptr(6),
				ptr(4), ptr(4),
				ptr(6),
				ptr(5), ptr(5), ptr(5), ptr(5),
				ptr(6),
				ptr(6), nil, nil, nil,
				nil,
				nil, nil, nil,
				nil,
				nil, nil, nil, nil,
				nil, nil,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day9.CompactBlocks(tt.uncompacted)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestCompactFiles(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		uncompacted []*int

		// outputs
		expected []*int
	}{
		"day's simplified example": {
			uncompacted: []*int{
				ptr(0),
				nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil, nil,
				ptr(2), ptr(2), ptr(2), ptr(2), ptr(2),
			},
			expected: []*int{
				ptr(0),
				nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil, nil,
				ptr(2), ptr(2), ptr(2), ptr(2), ptr(2),
			},
		},
		"day's full example": {
			uncompacted: []*int{
				ptr(0), ptr(0),
				nil, nil, nil,
				ptr(1), ptr(1), ptr(1),
				nil, nil, nil,
				ptr(2),
				nil, nil, nil,
				ptr(3), ptr(3), ptr(3),
				nil,
				ptr(4), ptr(4),
				nil,
				ptr(5), ptr(5), ptr(5), ptr(5),
				nil,
				ptr(6), ptr(6), ptr(6), ptr(6),
				nil,
				ptr(7), ptr(7), ptr(7),
				nil,
				ptr(8), ptr(8), ptr(8), ptr(8),
				ptr(9), ptr(9),
			},
			expected: []*int{
				ptr(0), ptr(0),
				ptr(9), ptr(9), ptr(2),
				ptr(1), ptr(1), ptr(1),
				ptr(7), ptr(7), ptr(7),
				nil,
				ptr(4), ptr(4), nil,
				ptr(3), ptr(3), ptr(3),
				nil,
				nil, nil,
				nil,
				ptr(5), ptr(5), ptr(5), ptr(5),
				nil,
				ptr(6), ptr(6), ptr(6), ptr(6),
				nil,
				nil, nil, nil,
				nil,
				ptr(8), ptr(8), ptr(8), ptr(8),
				nil, nil,
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day9.CompactFiles(tt.uncompacted)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestChecksum(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		blocks []*int

		// outputs
		expected int
	}{
		"day's simplified example": {
			blocks: []*int{
				ptr(0),
				ptr(2), ptr(2),
				ptr(1), ptr(1), ptr(1),
				ptr(2), ptr(2), ptr(2), nil,
				nil, nil, nil, nil, nil,
			},
			expected: 60,
		},
		"day's full example": {
			blocks: []*int{
				ptr(0), ptr(0),
				ptr(9), ptr(9), ptr(8),
				ptr(1), ptr(1), ptr(1),
				ptr(8), ptr(8), ptr(8),
				ptr(2),
				ptr(7), ptr(7), ptr(7),
				ptr(3), ptr(3), ptr(3),
				ptr(6),
				ptr(4), ptr(4),
				ptr(6),
				ptr(5), ptr(5), ptr(5), ptr(5),
				ptr(6),
				ptr(6), nil, nil, nil,
				nil,
				nil, nil, nil,
				nil,
				nil, nil, nil, nil,
				nil, nil,
			},
			expected: 1928,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day9.Checksum(tt.blocks)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func ptr[T any](v T) *T {
	return &v
}
