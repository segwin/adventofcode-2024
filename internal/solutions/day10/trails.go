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
	// Position of this trailhead.
	Position map2d.Position
	// Score is the number of unique summits reachable from this trailhead.
	Score int
	// Rating is the number of unique trails that start at this trailhead.
	Rating int
}

// TrailStep is a step along a trail. It contains a reference to all subsequent trail steps.
type TrailStep struct {
	Position map2d.Position
	Height   int
	Next     []TrailStep
}

// CountSummits returns the number of distinct summits reachable from this step.
func (s *TrailStep) CountSummits() int {
	return len(s.uniqueSummits())
}

// CountTrails returns the number of distinct trails that start at this step.
func (s *TrailStep) CountTrails() int {
	numTrails := 0
	for _, next := range s.Next {
		if next.Height == MaxHeight {
			numTrails++ // found summit
		} else {
			numTrails += next.CountTrails()
		}
	}
	return numTrails
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
