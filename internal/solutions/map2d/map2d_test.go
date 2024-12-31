package map2d_test

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
)

func TestMap_Get(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		m   map2d.Map[int]
		pos map2d.Position

		// outputs
		expected   int
		expectedOK bool
	}{
		"ok: map start": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 0, Y: 0},
			expected:   3,
			expectedOK: true,
		},
		"ok: map end": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 2, Y: 1},
			expected:   3,
			expectedOK: true,
		},
		"ok: mid-map": {
			m: [][]int{
				{3, 2, 1},
				{9, 8, 7},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 1, Y: 1},
			expected:   8,
			expectedOK: true,
		},

		"not ok: empty map": {
			m:          nil,
			pos:        map2d.Position{X: 0, Y: 0},
			expectedOK: false,
		},
		"not ok: X < map": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: -1, Y: 0},
			expectedOK: false,
		},
		"not ok: X > map": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 3, Y: 0},
			expectedOK: false,
		},
		"not ok: Y < map": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 0, Y: -1},
			expectedOK: false,
		},
		"not ok: Y > map": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:        map2d.Position{X: 0, Y: 2},
			expectedOK: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, ok := tt.m.Get(tt.pos)
			assert.Equal(t, tt.expectedOK, ok)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestMap_With(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		m        map2d.Map[int]
		pos      map2d.Position
		newValue int

		// outputs
		expected map2d.Map[int]
	}{
		"map start": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:      map2d.Position{X: 0, Y: 0},
			newValue: 9,
			expected: [][]int{
				{9, 2, 1},
				{1, 2, 3},
			},
		},
		"map end": {
			m: [][]int{
				{3, 2, 1},
				{1, 2, 3},
			},
			pos:      map2d.Position{X: 2, Y: 1},
			newValue: 9,
			expected: [][]int{
				{3, 2, 1},
				{1, 2, 9},
			},
		},
		"mid-map": {
			m: [][]int{
				{3, 2, 1},
				{9, 8, 7},
				{1, 2, 3},
			},
			pos:      map2d.Position{X: 1, Y: 1},
			newValue: 9,
			expected: [][]int{
				{3, 2, 1},
				{9, 9, 7},
				{1, 2, 3},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.m.With(tt.pos, tt.newValue)
			assert.Empty(t, cmp.Diff(tt.expected, got))
			assert.NotEmpty(t, cmp.Diff(got, tt.m)) // original map must not be altered
		})
	}
}

func TestPosition_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		// inputs
		p1, p2 map2d.Position

		// outputs
		expected map2d.Position
	}{
		{
			p1:       map2d.Position{X: 0, Y: 0},
			p2:       map2d.Position{X: 0, Y: 0},
			expected: map2d.Position{X: 0, Y: 0},
		},
		{
			p1:       map2d.Position{X: 1, Y: 2},
			p2:       map2d.Position{X: 0, Y: 0},
			expected: map2d.Position{X: 1, Y: 2},
		},
		{
			p1:       map2d.Position{X: 0, Y: 0},
			p2:       map2d.Position{X: 1, Y: 2},
			expected: map2d.Position{X: -1, Y: -2},
		},
		{
			p1:       map2d.Position{X: 3, Y: 2},
			p2:       map2d.Position{X: 1, Y: 4},
			expected: map2d.Position{X: 2, Y: -2},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got := tt.p1.Sub(tt.p2)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestPosition_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		// inputs
		p1, p2 map2d.Position

		// outputs
		expected map2d.Position
	}{
		{
			p1:       map2d.Position{X: 0, Y: 0},
			p2:       map2d.Position{X: 0, Y: 0},
			expected: map2d.Position{X: 0, Y: 0},
		},
		{
			p1:       map2d.Position{X: 1, Y: 2},
			p2:       map2d.Position{X: 0, Y: 0},
			expected: map2d.Position{X: 1, Y: 2},
		},
		{
			p1:       map2d.Position{X: 0, Y: 0},
			p2:       map2d.Position{X: 1, Y: 2},
			expected: map2d.Position{X: 1, Y: 2},
		},
		{
			p1:       map2d.Position{X: 3, Y: 2},
			p2:       map2d.Position{X: 1, Y: 4},
			expected: map2d.Position{X: 4, Y: 6},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got := tt.p1.Add(tt.p2)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}
