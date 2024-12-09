package solutions_test

import (
	"context"
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions"
	"github.com/stretchr/testify/require"
)

func TestRunOne(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		day int

		// outputs
		expectedErr error
	}{
		"error: day < 1": {
			day:         0,
			expectedErr: solutions.ErrInvalidDay,
		},
		"error: day > 25": {
			day:         26,
			expectedErr: solutions.ErrInvalidDay,
		},
		"ok: valid day": {
			day: 1,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			err := solutions.RunOne(context.Background(), tt.day)
			require.ErrorIs(t, err, tt.expectedErr)
		})
	}
}
