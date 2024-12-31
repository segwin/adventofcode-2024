package day10

import (
	"github.com/segwin/adventofcode-2024/internal/solutions/map2d"
)

type TopographicMap = map2d.Map[int]

const (
	MinHeight = 0
	MaxHeight = 9
)

// Trailhead is the start of a trail. It always has height = 0.
type Trailhead struct {
	Position map2d.Position
	Score    int
}

// TrailStep is a step along a trail. It contains a reference to all subsequent trail steps.
type TrailStep struct {
	Position map2d.Position
	Height   int
	Next     []TrailStep
}

// CountTrails returns the number of trails that start at this step.
func (s *TrailStep) CountTrails() int {
	return len(s.uniqueSummits())
}

func (s *TrailStep) uniqueSummits() map[map2d.Position]struct{} {
	summits := map[map2d.Position]struct{}{}
	for _, next := range s.Next {
		if next.Height == MaxHeight {
			// found summit
			summits[next.Position] = struct{}{}
			continue
		}
		// traverse next steps & add collected summits
		for pos := range next.uniqueSummits() {
			summits[pos] = struct{}{}
		}
	}
	return summits
}
