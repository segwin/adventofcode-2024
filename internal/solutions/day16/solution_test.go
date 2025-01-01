package day16_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day16"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day16.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}
