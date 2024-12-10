package day2_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSafeReports(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		// inputs
		reports          []day2.Report
		problemDampening bool

		// outputs
		expected    int
		expectedErr error
	}{
		"ok: day's example, no problem dampening": {
			reports: []day2.Report{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			problemDampening: false,
			expected:         2,
		},
		"ok: day's example, with problem dampening": {
			reports: []day2.Report{
				{7, 6, 4, 2, 1},
				{1, 2, 7, 8, 9},
				{9, 7, 6, 2, 1},
				{1, 3, 2, 4, 5},
				{8, 6, 4, 4, 1},
				{1, 3, 6, 7, 9},
			},
			problemDampening: true,
			expected:         4,
		},
		"ok: problem dampening where removal is not on first unsafe index": {
			reports: []day2.Report{
				{4, 5, 4, 7},
			},
			problemDampening: true,
			expected:         1,
		},
		"ok: report with <2 elements are always safe": {
			reports: []day2.Report{
				{},
				{123},
			},
			problemDampening: true,
			expected:         2,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := day2.SafeReports(tt.reports, tt.problemDampening)
			require.ErrorIs(t, err, tt.expectedErr)
			assert.Equal(t, tt.expected, got)
		})
	}
}
