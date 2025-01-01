package day17_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day17"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day17.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}
