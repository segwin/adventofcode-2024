package map2d_test

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
	"github.com/stretchr/testify/assert"
)

func TestPosition_Sub(t *testing.T) {
	t.Parallel()

	tests := []struct {
		// inputs
		p    map2d.Position
		x, y int

		// outputs
		expected map2d.Distance
	}{
		{
			p: map2d.Position{X: 0, Y: 0},
			x: 0, y: 0,
			expected: map2d.Distance{X: 0, Y: 0},
		},
		{
			p: map2d.Position{X: 1, Y: 2},
			x: 0, y: 0,
			expected: map2d.Distance{X: 1, Y: 2},
		},
		{
			p: map2d.Position{X: 0, Y: 0},
			x: 1, y: 2,
			expected: map2d.Distance{X: -1, Y: -2},
		},
		{
			p: map2d.Position{X: 3, Y: 2},
			x: 1, y: 4,
			expected: map2d.Distance{X: 2, Y: -2},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got := tt.p.Sub(tt.x, tt.y)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestPosition_Add(t *testing.T) {
	t.Parallel()

	tests := []struct {
		// inputs
		p    map2d.Position
		x, y int

		// outputs
		expected map2d.Position
	}{
		{
			p: map2d.Position{X: 0, Y: 0},
			x: 0, y: 0,
			expected: map2d.Position{X: 0, Y: 0},
		},
		{
			p: map2d.Position{X: 1, Y: 2},
			x: 0, y: 0,
			expected: map2d.Position{X: 1, Y: 2},
		},
		{
			p: map2d.Position{X: 0, Y: 0},
			x: 1, y: 2,
			expected: map2d.Position{X: 1, Y: 2},
		},
		{
			p: map2d.Position{X: 3, Y: 2},
			x: 1, y: 4,
			expected: map2d.Position{X: 4, Y: 6},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			got := tt.p.Add(tt.x, tt.y)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestPosition_Move(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		p      map2d.Position
		d      map2d.Direction
		amount int

		// outputs
		expected map2d.Position
	}{
		"move north by 1": {
			p: map2d.Position{}, d: map2d.North(), amount: 1,
			expected: map2d.Position{X: 0, Y: -1},
		},
		"move north by 10": {
			p: map2d.Position{}, d: map2d.North(), amount: 10,
			expected: map2d.Position{X: 0, Y: -10},
		},
		"move east by 1": {
			p: map2d.Position{}, d: map2d.East(), amount: 1,
			expected: map2d.Position{X: 1, Y: 0},
		},
		"move east by 10": {
			p: map2d.Position{}, d: map2d.East(), amount: 10,
			expected: map2d.Position{X: 10, Y: 0},
		},
		"move south by 1": {
			p: map2d.Position{}, d: map2d.South(), amount: 1,
			expected: map2d.Position{X: 0, Y: 1},
		},
		"move south by 10": {
			p: map2d.Position{}, d: map2d.South(), amount: 10,
			expected: map2d.Position{X: 0, Y: 10},
		},
		"move west by 1": {
			p: map2d.Position{}, d: map2d.West(), amount: 1,
			expected: map2d.Position{X: -1, Y: 0},
		},
		"move west by 10": {
			p: map2d.Position{}, d: map2d.West(), amount: 10,
			expected: map2d.Position{X: -10, Y: 0},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.p.Move(tt.d, tt.amount)
			assert.Empty(t, cmp.Diff(tt.expected, got))
		})
	}
}

func TestPosition_ProjectOnto(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		p map2d.Position
		d map2d.Direction

		// outputs
		expected int
	}{
		"north, y = 0 => 0": {
			p:        map2d.Position{X: 1, Y: 0},
			d:        map2d.North(),
			expected: 0,
		},
		"north, y = 4 => -4": {
			p:        map2d.Position{X: 1, Y: 4},
			d:        map2d.North(),
			expected: -4,
		},
		"east, x = 0 => 0": {
			p:        map2d.Position{X: 0, Y: 1},
			d:        map2d.East(),
			expected: 0,
		},
		"east, x = 4 => 4": {
			p:        map2d.Position{X: 4, Y: 1},
			d:        map2d.East(),
			expected: 4,
		},
		"south, y = 0 => 0": {
			p:        map2d.Position{X: 1, Y: 0},
			d:        map2d.South(),
			expected: 0,
		},
		"south, y = 4 => 4": {
			p:        map2d.Position{X: 1, Y: 4},
			d:        map2d.South(),
			expected: 4,
		},
		"west, x = 0 => 0": {
			p:        map2d.Position{X: 0, Y: 1},
			d:        map2d.West(),
			expected: 0,
		},
		"west, x = 4 => -4": {
			p:        map2d.Position{X: 4, Y: 1},
			d:        map2d.West(),
			expected: -4,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tt.p.ProjectOnto(tt.d)
			assert.Equal(t, tt.expected, got)
		})
	}
}
