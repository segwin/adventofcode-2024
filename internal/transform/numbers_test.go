package transform_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/transform"
	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		in int

		// outputs
		expected int
	}{
		"0 => 0":   {in: 0, expected: 0},
		"-1 => +1": {in: -1, expected: +1},
		"+1 => +1": {in: +1, expected: +1},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, transform.Abs(tt.in))
		})
	}
}
