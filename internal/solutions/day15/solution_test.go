package day15_test

import (
	"testing"

	"github.com/segwin/adventofcode-2024/internal/solutions/day15"
	"github.com/stretchr/testify/require"
)

func TestSolution(t *testing.T) {
	t.Parallel()

	s, err := day15.BuildSolution()
	require.NoError(t, err)
	require.NoError(t, s.RunToConsole())
}
