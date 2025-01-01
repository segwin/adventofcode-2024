package day8_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day8.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}

func TestUniqueAntinodes(t *testing.T) {
	t.Parallel()

	exampleInput := day8.CityMap{
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '0', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '0', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', 'A', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', 'A', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}

	tests := map[string]struct {
		// inputs
		cityMap          day8.CityMap
		includeHarmonics bool

		// outputs
		expected int
	}{
		"ok: day's example, no harmonics": {
			cityMap:          exampleInput,
			includeHarmonics: false,
			expected:         14,
		},
		"ok: day's example, with harmonics": {
			cityMap:          exampleInput,
			includeHarmonics: true,
			expected:         34,
		},
		"ok: extra example, with harmonics": {
			cityMap: day8.CityMap{
				{'T', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', 'T', '.', '.', '.', '.', '.', '.'},
				{'.', 'T', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
			},
			includeHarmonics: true,
			expected:         9,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := day8.UniqueAntinodeLocations(tt.cityMap, tt.includeHarmonics)
			assert.Equal(t, tt.expected, got)
		})
	}
}
